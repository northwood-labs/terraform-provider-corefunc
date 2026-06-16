data "corefunc_json_jq" "jq" {
  # json = file("${path.module}/input.json")
  json = data.kubernetes_pod.v1_pods_json.json

  # query     = ".items[] | {name: .metadata.name, namespace: .metadata.namespace}"
  query = file("${path.module}/query.jq")
}

# provider::corefunc::json_jq(json, query)

data "corefunc_jsonschema_validate" "js" {
  json       = data.corefunc_json_jq.jq.result_json
  jsonschema = file("${path.module}/schema.json")
}

data "corefunc_jsonschema_validate" "js" {
  json = jsonencode({
    name      = "example-pod"
    namespace = "default"
  })
  jsonschema = file("${path.module}/schema.json")
}

# provider::corefunc::jsonschema_validate(jsonschema, json)
