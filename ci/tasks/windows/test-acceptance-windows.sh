#!/bin/bash -eux

set -eu -o pipefail

ROOT_DIR=$PWD
BBL_STATE_DIR=$ROOT_DIR/envs/$ENV_NAME

pushd "${BBL_STATE_DIR}"
  source .envrc
popd

export BOSH_DEPLOYMENT=bosh-dns-windows-acceptance

# Need to delete the bosh-dns runtime config because bbl uses a hard-coded
# bosh-deployment which specifies a bosh-dns version that may conflict with the
# one we are trying to test.
bosh delete-config --type=runtime --name=bosh-dns -n

bosh -n deploy \
  $ROOT_DIR/bosh-dns-release/src/bosh-dns/test_yml_assets/manifests/windows-acceptance-manifest.yml \
  -v deployment_name="$BOSH_DEPLOYMENT" \
  -v windows_stemcell=$WINDOWS_OS_VERSION \
  --vars-store dns-creds.yml \
  -o $ROOT_DIR/bosh-dns-release/src/bosh-dns/test_yml_assets/ops/enable-health-manifest-ops.yml \
  -v health_server_port=2345

bosh run-errand acceptance-tests-windows
