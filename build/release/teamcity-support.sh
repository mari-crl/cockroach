# Common helpers for teamcity-*.sh scripts.

# root is the absolute path to the root directory of the repository.
root="$(dirname $(dirname $(cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )))"
source "$root/build/teamcity-common-support.sh"

declare -a platform_names=( "amd64::amd64" "aarch64::arm64" )
declare -a tarball_archs=( "${platform_names[@]%%::*}" ) 
declare -a docker_archs=( "${platform_names[@]##*::}" )

function tarball_arch_from_platform_name() {
  echo "${1%%::*}"
}

function docker_arch_from_platform_name() {
  echo "${1##*::}"
}

remove_files_on_exit() {
  rm -rf ~/.docker
  common_support_remove_files_on_exit
}
trap remove_files_on_exit EXIT

tc_start_block() {
  echo "##teamcity[blockOpened name='$1']"
}

tc_end_block() {
  echo "##teamcity[blockClosed name='$1']"
}

docker_login_with_google() {
  # https://cloud.google.com/container-registry/docs/advanced-authentication#json-key
  echo "${google_credentials}" | docker login -u _json_key --password-stdin "https://${gcr_hostname}"
}

docker_login() {
  echo "${DOCKER_ACCESS_TOKEN}" | docker login --username "${DOCKER_ID}" --password-stdin
}

configure_docker_creds() {
  # Work around headless d-bus problem by forcing docker to use
  # plain-text credentials for dockerhub.
  # https://github.com/docker/docker-credential-helpers/issues/105#issuecomment-420480401
  mkdir -p ~/.docker
  cat << EOF > ~/.docker/config.json
{
  "credsStore" : "",
  "auths": {
    "https://index.docker.io/v1/" : {
    }
  }
}
EOF
}

docker_login_with_redhat() {
  echo "${REDHAT_REGISTRY_KEY}" | docker login --username unused --password-stdin $rhel_registry
}
