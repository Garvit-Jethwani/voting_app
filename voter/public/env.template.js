(function (window) {
  window["env"] = window["env"] || {};

  // Environment variables
  var ballotEndpointEnvVar = `$REACT_APP_BALLOT_ENDPOINT`;
  var ecServerEndpoint = `$REACT_APP_EC_SERVER_ENDPOINT`;
  var ldClientId = `$REACT_APP_LD_CLIENT_ID`;
  window["env"]["ballotEndpoint"] = ballotEndpointEnvVar.replace(/['"]/gi, "");
  window["env"]["ecServerEndpoint"] = ecServerEndpoint.replace(/['"]/gi, "");
  window["env"]["ld_client_id"] = ldClientId.replace(/['"]/gi, "");
})(this);
