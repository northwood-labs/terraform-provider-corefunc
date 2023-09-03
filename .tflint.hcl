# .TFLINT.HCL CONFIG
# Generated from generate-tflint-config
# DO NOT EDIT!

plugin "aws" {
  enabled    = true
  deep_check = true
  version    = "0.20.0" # This should be the same version as in `generate-tflint-config/Makefile`.
  source     = "github.com/terraform-linters/tflint-ruleset-aws"
}

config {
  module = true
}

#-------------------------------------------------------------------------------
# https://github.com/terraform-linters/tflint/blob/master/docs/rules/README.md

rule "terraform_comment_syntax" {
  enabled = true
}

rule "terraform_deprecated_index" {
  enabled = true
}

rule "terraform_deprecated_interpolation" {
  enabled = true
}

rule "terraform_documented_outputs" {
  enabled = true
}

rule "terraform_documented_variables" {
  enabled = true
}

rule "terraform_module_pinned_source" {
  enabled = true
  style   = "semver"
}

rule "terraform_naming_convention" {
  enabled = true
  format  = "snake_case"
}

rule "terraform_required_providers" {
  enabled = true
}

rule "terraform_required_version" {
  enabled = true
}

rule "terraform_standard_module_structure" {
  enabled = false
}

rule "terraform_typed_variables" {
  enabled = true
}

rule "terraform_unused_declarations" {
  enabled = false
}

rule "terraform_unused_required_providers" {
  enabled = true
}

rule "terraform_workspace_remote" {
  enabled = false
}

#-------------------------------------------------------------------------------
# Possible Errors
# https://github.com/terraform-linters/tflint-ruleset-aws/blob/master/docs/rules/README.md

rule "aws_alb_invalid_security_group" {
  enabled = true
}

rule "aws_alb_invalid_subnet" {
  enabled = true
}

rule "aws_api_gateway_model_invalid_name" {
  enabled = true
}

rule "aws_db_instance_invalid_db_subnet_group" {
  enabled = true
}

rule "aws_db_instance_invalid_engine" {
  enabled = true
}

rule "aws_db_instance_invalid_option_group" {
  enabled = true
}

rule "aws_db_instance_invalid_parameter_group" {
  enabled = true
}

rule "aws_db_instance_invalid_type" {
  enabled = true
}

rule "aws_db_instance_invalid_vpc_security_group" {
  enabled = true
}

rule "aws_elasticache_cluster_invalid_parameter_group" {
  enabled = true
}

rule "aws_elasticache_cluster_invalid_security_group" {
  enabled = true
}

rule "aws_elasticache_cluster_invalid_subnet_group" {
  enabled = true
}

rule "aws_elasticache_cluster_invalid_type" {
  enabled = true
}

rule "aws_elasticache_replication_group_invalid_type" {
  enabled = true
}

rule "aws_elb_invalid_instance" {
  enabled = true
}

rule "aws_elb_invalid_security_group" {
  enabled = true
}

rule "aws_elb_invalid_subnet" {
  enabled = true
}

rule "aws_instance_invalid_ami" {
  enabled = true
}

rule "aws_instance_invalid_iam_profile" {
  enabled = true
}

rule "aws_instance_invalid_key_name" {
  enabled = true
}

rule "aws_instance_invalid_subnet" {
  enabled = true
}

rule "aws_instance_invalid_vpc_security_group" {
  enabled = true
}

rule "aws_launch_configuration_invalid_iam_profile" {
  enabled = true
}

rule "aws_launch_configuration_invalid_image_id" {
  enabled = true
}

rule "aws_route_invalid_egress_only_gateway" {
  enabled = true
}

rule "aws_route_invalid_gateway" {
  enabled = true
}

rule "aws_route_invalid_instance" {
  enabled = true
}

rule "aws_route_invalid_nat_gateway" {
  enabled = true
}

rule "aws_route_invalid_network_interface" {
  enabled = true
}

rule "aws_route_invalid_route_table" {
  enabled = true
}

rule "aws_route_invalid_vpc_peering_connection" {
  enabled = true
}

rule "aws_route_not_specified_target" {
  enabled = true
}

rule "aws_route_specified_multiple_targets" {
  enabled = true
}

#-------------------------------------------------------------------------------
# Best Practices/Naming Conventions
# https://github.com/terraform-linters/tflint-ruleset-aws/blob/master/docs/rules/README.md

rule "aws_db_instance_previous_type" {
  enabled = true
}

rule "aws_db_instance_default_parameter_group" {
  enabled = true
}

rule "aws_elasticache_cluster_previous_type" {
  enabled = true
}

rule "aws_elasticache_cluster_default_parameter_group" {
  enabled = true
}

rule "aws_elasticache_replication_group_previous_type" {
  enabled = true
}

rule "aws_elasticache_replication_group_default_parameter_group" {
  enabled = true
}

rule "aws_instance_previous_type" {
  enabled = true
}

rule "aws_iam_policy_document_gov_friendly_arns" {
  enabled = false
}

rule "aws_iam_policy_gov_friendly_arns" {
  enabled = false
}

rule "aws_iam_role_policy_gov_friendly_arns" {
  enabled = false
}

rule "aws_resource_missing_tags" {
  enabled = false
}

rule "aws_s3_bucket_name" {
  enabled = true
}

#-------------------------------------------------------------------------------
# SDK-based Validations
# https://github.com/terraform-linters/tflint-ruleset-aws/blob/master/docs/rules/README.md
rule "aws_accessanalyzer_analyzer_invalid_analyzer_name" {
  enabled = true
}

rule "aws_accessanalyzer_analyzer_invalid_type" {
  enabled = true
}

rule "aws_account_alternate_contact_invalid_account_id" {
  enabled = true
}

rule "aws_account_alternate_contact_invalid_alternate_contact_type" {
  enabled = true
}

rule "aws_account_alternate_contact_invalid_email_address" {
  enabled = true
}

rule "aws_account_alternate_contact_invalid_name" {
  enabled = true
}

rule "aws_account_alternate_contact_invalid_phone_number" {
  enabled = true
}

rule "aws_account_alternate_contact_invalid_title" {
  enabled = true
}

rule "aws_acm_certificate_invalid_certificate_authority_arn" {
  enabled = true
}

rule "aws_acm_certificate_invalid_certificate_body" {
  enabled = true
}

rule "aws_acm_certificate_invalid_certificate_chain" {
  enabled = true
}

rule "aws_acm_certificate_invalid_private_key" {
  enabled = true
}

rule "aws_acm_certificate_validation_invalid_certificate_arn" {
  enabled = true
}

rule "aws_acmpca_certificate_authority_certificate_invalid_certificate_authority_arn" {
  enabled = true
}

rule "aws_acmpca_certificate_authority_invalid_type" {
  enabled = true
}

rule "aws_acmpca_certificate_invalid_certificate_authority_arn" {
  enabled = true
}

rule "aws_acmpca_certificate_invalid_signing_algorithm" {
  enabled = true
}

rule "aws_alb_invalid_ip_address_type" {
  enabled = true
}

rule "aws_alb_invalid_load_balancer_type" {
  enabled = true
}

rule "aws_alb_listener_invalid_protocol" {
  enabled = true
}

rule "aws_alb_target_group_invalid_protocol" {
  enabled = true
}

rule "aws_alb_target_group_invalid_target_type" {
  enabled = true
}

rule "aws_ami_invalid_architecture" {
  enabled = true
}

rule "aws_amplify_app_invalid_access_token" {
  enabled = true
}

rule "aws_amplify_app_invalid_basic_auth_credentials" {
  enabled = true
}

rule "aws_amplify_app_invalid_build_spec" {
  enabled = true
}

rule "aws_amplify_app_invalid_description" {
  enabled = true
}

rule "aws_amplify_app_invalid_iam_service_role_arn" {
  enabled = true
}

rule "aws_amplify_app_invalid_name" {
  enabled = true
}

rule "aws_amplify_app_invalid_oauth_token" {
  enabled = true
}

rule "aws_amplify_app_invalid_platform" {
  enabled = true
}

rule "aws_amplify_app_invalid_repository" {
  enabled = true
}

rule "aws_amplify_backend_environment_invalid_app_id" {
  enabled = true
}

rule "aws_amplify_backend_environment_invalid_deployment_artifacts" {
  enabled = true
}

rule "aws_amplify_backend_environment_invalid_environment_name" {
  enabled = true
}

rule "aws_amplify_backend_environment_invalid_stack_name" {
  enabled = true
}

rule "aws_amplify_branch_invalid_app_id" {
  enabled = true
}

rule "aws_amplify_branch_invalid_backend_environment_arn" {
  enabled = true
}

rule "aws_amplify_branch_invalid_basic_auth_credentials" {
  enabled = true
}

rule "aws_amplify_branch_invalid_branch_name" {
  enabled = true
}

rule "aws_amplify_branch_invalid_description" {
  enabled = true
}

rule "aws_amplify_branch_invalid_display_name" {
  enabled = true
}

rule "aws_amplify_branch_invalid_framework" {
  enabled = true
}

rule "aws_amplify_branch_invalid_pull_request_environment_name" {
  enabled = true
}

rule "aws_amplify_branch_invalid_stage" {
  enabled = true
}

rule "aws_amplify_branch_invalid_ttl" {
  enabled = true
}

rule "aws_amplify_domain_association_invalid_app_id" {
  enabled = true
}

rule "aws_amplify_webhook_invalid_app_id" {
  enabled = true
}

rule "aws_amplify_webhook_invalid_branch_name" {
  enabled = true
}

rule "aws_amplify_webhook_invalid_description" {
  enabled = true
}

rule "aws_api_gateway_authorizer_invalid_type" {
  enabled = true
}

rule "aws_api_gateway_domain_name_invalid_security_policy" {
  enabled = true
}

rule "aws_api_gateway_gateway_response_invalid_response_type" {
  enabled = true
}

rule "aws_api_gateway_gateway_response_invalid_status_code" {
  enabled = true
}

rule "aws_api_gateway_integration_invalid_connection_type" {
  enabled = true
}

rule "aws_api_gateway_integration_invalid_content_handling" {
  enabled = true
}

rule "aws_api_gateway_integration_invalid_type" {
  enabled = true
}

rule "aws_api_gateway_integration_response_invalid_content_handling" {
  enabled = true
}

rule "aws_api_gateway_integration_response_invalid_status_code" {
  enabled = true
}

rule "aws_api_gateway_method_response_invalid_status_code" {
  enabled = true
}

rule "aws_api_gateway_rest_api_invalid_api_key_source" {
  enabled = true
}

rule "aws_api_gateway_stage_invalid_cache_cluster_size" {
  enabled = true
}

rule "aws_apigatewayv2_api_invalid_protocol_type" {
  enabled = true
}

rule "aws_apigatewayv2_authorizer_invalid_authorizer_type" {
  enabled = true
}

rule "aws_apigatewayv2_integration_invalid_connection_type" {
  enabled = true
}

rule "aws_apigatewayv2_integration_invalid_content_handling_strategy" {
  enabled = true
}

rule "aws_apigatewayv2_integration_invalid_integration_type" {
  enabled = true
}

rule "aws_apigatewayv2_integration_invalid_passthrough_behavior" {
  enabled = true
}

rule "aws_apigatewayv2_integration_response_invalid_content_handling_strategy" {
  enabled = true
}

rule "aws_apigatewayv2_route_invalid_authorization_type" {
  enabled = true
}

rule "aws_appautoscaling_policy_invalid_policy_type" {
  enabled = true
}

rule "aws_appautoscaling_policy_invalid_scalable_dimension" {
  enabled = true
}

rule "aws_appautoscaling_policy_invalid_service_namespace" {
  enabled = true
}

rule "aws_appautoscaling_scheduled_action_invalid_scalable_dimension" {
  enabled = true
}

rule "aws_appautoscaling_scheduled_action_invalid_service_namespace" {
  enabled = true
}

rule "aws_appautoscaling_target_invalid_scalable_dimension" {
  enabled = true
}

rule "aws_appautoscaling_target_invalid_service_namespace" {
  enabled = true
}

rule "aws_appconfig_application_invalid_description" {
  enabled = true
}

rule "aws_appconfig_application_invalid_name" {
  enabled = true
}

rule "aws_appconfig_configuration_profile_invalid_application_id" {
  enabled = true
}

rule "aws_appconfig_configuration_profile_invalid_description" {
  enabled = true
}

rule "aws_appconfig_configuration_profile_invalid_location_uri" {
  enabled = true
}

rule "aws_appconfig_configuration_profile_invalid_name" {
  enabled = true
}

rule "aws_appconfig_configuration_profile_invalid_retrieval_role_arn" {
  enabled = true
}

rule "aws_appconfig_deployment_invalid_application_id" {
  enabled = true
}

rule "aws_appconfig_deployment_invalid_configuration_profile_id" {
  enabled = true
}

rule "aws_appconfig_deployment_invalid_configuration_version" {
  enabled = true
}

rule "aws_appconfig_deployment_invalid_deployment_strategy_id" {
  enabled = true
}

rule "aws_appconfig_deployment_invalid_description" {
  enabled = true
}

rule "aws_appconfig_deployment_invalid_environment_id" {
  enabled = true
}

rule "aws_appconfig_deployment_strategy_invalid_description" {
  enabled = true
}

rule "aws_appconfig_deployment_strategy_invalid_growth_type" {
  enabled = true
}

rule "aws_appconfig_deployment_strategy_invalid_name" {
  enabled = true
}

rule "aws_appconfig_deployment_strategy_invalid_replicate_to" {
  enabled = true
}

rule "aws_appconfig_environment_invalid_application_id" {
  enabled = true
}

rule "aws_appconfig_environment_invalid_description" {
  enabled = true
}

rule "aws_appconfig_environment_invalid_name" {
  enabled = true
}

rule "aws_appconfig_hosted_configuration_version_invalid_application_id" {
  enabled = true
}

rule "aws_appconfig_hosted_configuration_version_invalid_configuration_profile_id" {
  enabled = true
}

rule "aws_appconfig_hosted_configuration_version_invalid_content_type" {
  enabled = true
}

rule "aws_appconfig_hosted_configuration_version_invalid_description" {
  enabled = true
}

rule "aws_appmesh_gateway_route_invalid_mesh_name" {
  enabled = true
}

rule "aws_appmesh_gateway_route_invalid_mesh_owner" {
  enabled = true
}

rule "aws_appmesh_gateway_route_invalid_name" {
  enabled = true
}

rule "aws_appmesh_gateway_route_invalid_virtual_gateway_name" {
  enabled = true
}

rule "aws_appmesh_mesh_invalid_name" {
  enabled = true
}

rule "aws_appmesh_route_invalid_mesh_name" {
  enabled = true
}

rule "aws_appmesh_route_invalid_name" {
  enabled = true
}

rule "aws_appmesh_route_invalid_virtual_router_name" {
  enabled = true
}

rule "aws_appmesh_virtual_gateway_invalid_mesh_name" {
  enabled = true
}

rule "aws_appmesh_virtual_gateway_invalid_mesh_owner" {
  enabled = true
}

rule "aws_appmesh_virtual_gateway_invalid_name" {
  enabled = true
}

rule "aws_appmesh_virtual_node_invalid_mesh_name" {
  enabled = true
}

rule "aws_appmesh_virtual_node_invalid_name" {
  enabled = true
}

rule "aws_appmesh_virtual_router_invalid_mesh_name" {
  enabled = true
}

rule "aws_appmesh_virtual_router_invalid_name" {
  enabled = true
}

rule "aws_appmesh_virtual_service_invalid_mesh_name" {
  enabled = true
}

rule "aws_appmesh_virtual_service_invalid_name" {
  enabled = true
}

rule "aws_apprunner_auto_scaling_configuration_version_invalid_auto_scaling_configuration_name" {
  enabled = true
}

rule "aws_apprunner_connection_invalid_connection_name" {
  enabled = true
}

rule "aws_apprunner_connection_invalid_provider_type" {
  enabled = true
}

rule "aws_apprunner_custom_domain_association_invalid_domain_name" {
  enabled = true
}

rule "aws_apprunner_service_invalid_service_name" {
  enabled = true
}

rule "aws_appstream_fleet_invalid_description" {
  enabled = true
}

rule "aws_appstream_fleet_invalid_display_name" {
  enabled = true
}

rule "aws_appstream_fleet_invalid_fleet_type" {
  enabled = true
}

rule "aws_appstream_fleet_invalid_name" {
  enabled = true
}

rule "aws_appstream_fleet_invalid_stream_view" {
  enabled = true
}

rule "aws_appstream_image_builder_invalid_appstream_agent_version" {
  enabled = true
}

rule "aws_appstream_image_builder_invalid_description" {
  enabled = true
}

rule "aws_appstream_image_builder_invalid_display_name" {
  enabled = true
}

rule "aws_appstream_image_builder_invalid_name" {
  enabled = true
}

rule "aws_appstream_stack_invalid_description" {
  enabled = true
}

rule "aws_appstream_stack_invalid_display_name" {
  enabled = true
}

rule "aws_appstream_stack_invalid_feedback_url" {
  enabled = true
}

rule "aws_appstream_stack_invalid_redirect_url" {
  enabled = true
}

rule "aws_appstream_user_invalid_authentication_type" {
  enabled = true
}

rule "aws_appstream_user_invalid_first_name" {
  enabled = true
}

rule "aws_appstream_user_invalid_last_name" {
  enabled = true
}

rule "aws_appstream_user_invalid_user_name" {
  enabled = true
}

rule "aws_appstream_user_stack_association_invalid_authentication_type" {
  enabled = true
}

rule "aws_appstream_user_stack_association_invalid_user_name" {
  enabled = true
}

rule "aws_appsync_datasource_invalid_name" {
  enabled = true
}

rule "aws_appsync_datasource_invalid_type" {
  enabled = true
}

rule "aws_appsync_function_invalid_data_source" {
  enabled = true
}

rule "aws_appsync_function_invalid_name" {
  enabled = true
}

rule "aws_appsync_function_invalid_request_mapping_template" {
  enabled = true
}

rule "aws_appsync_function_invalid_response_mapping_template" {
  enabled = true
}

rule "aws_appsync_graphql_api_invalid_authentication_type" {
  enabled = true
}

rule "aws_appsync_resolver_invalid_data_source" {
  enabled = true
}

rule "aws_appsync_resolver_invalid_field" {
  enabled = true
}

rule "aws_appsync_resolver_invalid_request_template" {
  enabled = true
}

rule "aws_appsync_resolver_invalid_response_template" {
  enabled = true
}

rule "aws_appsync_resolver_invalid_type" {
  enabled = true
}

rule "aws_athena_database_invalid_name" {
  enabled = true
}

rule "aws_athena_named_query_invalid_database" {
  enabled = true
}

rule "aws_athena_named_query_invalid_description" {
  enabled = true
}

rule "aws_athena_named_query_invalid_name" {
  enabled = true
}

rule "aws_athena_named_query_invalid_query" {
  enabled = true
}

rule "aws_athena_workgroup_invalid_description" {
  enabled = true
}

rule "aws_athena_workgroup_invalid_name" {
  enabled = true
}

rule "aws_athena_workgroup_invalid_state" {
  enabled = true
}

rule "aws_backup_selection_invalid_name" {
  enabled = true
}

rule "aws_backup_vault_invalid_name" {
  enabled = true
}

rule "aws_backup_vault_lock_configuration_invalid_backup_vault_name" {
  enabled = true
}

rule "aws_backup_vault_notifications_invalid_backup_vault_name" {
  enabled = true
}

rule "aws_backup_vault_policy_invalid_backup_vault_name" {
  enabled = true
}

rule "aws_batch_compute_environment_invalid_state" {
  enabled = true
}

rule "aws_batch_compute_environment_invalid_type" {
  enabled = true
}

rule "aws_batch_job_definition_invalid_type" {
  enabled = true
}

rule "aws_batch_job_queue_invalid_state" {
  enabled = true
}

rule "aws_budgets_budget_invalid_account_id" {
  enabled = true
}

rule "aws_budgets_budget_invalid_budget_type" {
  enabled = true
}

rule "aws_budgets_budget_invalid_name" {
  enabled = true
}

rule "aws_budgets_budget_invalid_time_unit" {
  enabled = true
}

rule "aws_chime_voice_connector_group_invalid_name" {
  enabled = true
}

rule "aws_chime_voice_connector_invalid_aws_region" {
  enabled = true
}

rule "aws_chime_voice_connector_invalid_name" {
  enabled = true
}

rule "aws_chime_voice_connector_logging_invalid_voice_connector_id" {
  enabled = true
}

rule "aws_chime_voice_connector_origination_invalid_voice_connector_id" {
  enabled = true
}

rule "aws_chime_voice_connector_streaming_invalid_voice_connector_id" {
  enabled = true
}

rule "aws_chime_voice_connector_termination_credentials_invalid_voice_connector_id" {
  enabled = true
}

rule "aws_chime_voice_connector_termination_invalid_default_phone_number" {
  enabled = true
}

rule "aws_chime_voice_connector_termination_invalid_voice_connector_id" {
  enabled = true
}

rule "aws_cloud9_environment_ec2_invalid_description" {
  enabled = true
}

rule "aws_cloud9_environment_ec2_invalid_instance_type" {
  enabled = true
}

rule "aws_cloud9_environment_ec2_invalid_name" {
  enabled = true
}

rule "aws_cloud9_environment_ec2_invalid_owner_arn" {
  enabled = true
}

rule "aws_cloud9_environment_ec2_invalid_subnet_id" {
  enabled = true
}

rule "aws_cloudformation_stack_invalid_iam_role_arn" {
  enabled = true
}

rule "aws_cloudformation_stack_invalid_on_failure" {
  enabled = true
}

rule "aws_cloudformation_stack_invalid_policy_body" {
  enabled = true
}

rule "aws_cloudformation_stack_invalid_policy_url" {
  enabled = true
}

rule "aws_cloudformation_stack_invalid_template_url" {
  enabled = true
}

rule "aws_cloudformation_stack_set_instance_invalid_account_id" {
  enabled = true
}

rule "aws_cloudformation_stack_set_invalid_administration_role_arn" {
  enabled = true
}

rule "aws_cloudformation_stack_set_invalid_description" {
  enabled = true
}

rule "aws_cloudformation_stack_set_invalid_execution_role_name" {
  enabled = true
}

rule "aws_cloudformation_stack_set_invalid_template_url" {
  enabled = true
}

rule "aws_cloudfront_distribution_invalid_http_version" {
  enabled = true
}

rule "aws_cloudfront_distribution_invalid_price_class" {
  enabled = true
}

rule "aws_cloudfront_function_invalid_name" {
  enabled = true
}

rule "aws_cloudfront_function_invalid_runtime" {
  enabled = true
}

rule "aws_cloudhsm_v2_cluster_invalid_hsm_type" {
  enabled = true
}

rule "aws_cloudhsm_v2_cluster_invalid_source_backup_identifier" {
  enabled = true
}

rule "aws_cloudhsm_v2_hsm_invalid_availability_zone" {
  enabled = true
}

rule "aws_cloudhsm_v2_hsm_invalid_cluster_id" {
  enabled = true
}

rule "aws_cloudhsm_v2_hsm_invalid_ip_address" {
  enabled = true
}

rule "aws_cloudhsm_v2_hsm_invalid_subnet_id" {
  enabled = true
}

rule "aws_cloudwatch_event_api_destination_invalid_connection_arn" {
  enabled = true
}

rule "aws_cloudwatch_event_api_destination_invalid_description" {
  enabled = true
}

rule "aws_cloudwatch_event_api_destination_invalid_http_method" {
  enabled = true
}

rule "aws_cloudwatch_event_api_destination_invalid_invocation_endpoint" {
  enabled = true
}

rule "aws_cloudwatch_event_api_destination_invalid_name" {
  enabled = true
}

rule "aws_cloudwatch_event_archive_invalid_description" {
  enabled = true
}

rule "aws_cloudwatch_event_archive_invalid_event_source_arn" {
  enabled = true
}

rule "aws_cloudwatch_event_archive_invalid_name" {
  enabled = true
}

rule "aws_cloudwatch_event_bus_invalid_event_source_name" {
  enabled = true
}

rule "aws_cloudwatch_event_bus_invalid_name" {
  enabled = true
}

rule "aws_cloudwatch_event_bus_policy_invalid_event_bus_name" {
  enabled = true
}

rule "aws_cloudwatch_event_connection_invalid_authorization_type" {
  enabled = true
}

rule "aws_cloudwatch_event_connection_invalid_description" {
  enabled = true
}

rule "aws_cloudwatch_event_connection_invalid_name" {
  enabled = true
}

rule "aws_cloudwatch_event_permission_invalid_action" {
  enabled = true
}

rule "aws_cloudwatch_event_permission_invalid_principal" {
  enabled = true
}

rule "aws_cloudwatch_event_permission_invalid_statement_id" {
  enabled = true
}

rule "aws_cloudwatch_event_rule_invalid_description" {
  enabled = true
}

rule "aws_cloudwatch_event_rule_invalid_name" {
  enabled = true
}

rule "aws_cloudwatch_event_rule_invalid_role_arn" {
  enabled = true
}

rule "aws_cloudwatch_event_rule_invalid_schedule_expression" {
  enabled = true
}

rule "aws_cloudwatch_event_target_invalid_arn" {
  enabled = true
}

rule "aws_cloudwatch_event_target_invalid_input" {
  enabled = true
}

rule "aws_cloudwatch_event_target_invalid_input_path" {
  enabled = true
}

rule "aws_cloudwatch_event_target_invalid_role_arn" {
  enabled = true
}

rule "aws_cloudwatch_event_target_invalid_rule" {
  enabled = true
}

rule "aws_cloudwatch_event_target_invalid_target_id" {
  enabled = true
}

rule "aws_cloudwatch_log_destination_invalid_name" {
  enabled = true
}

rule "aws_cloudwatch_log_destination_policy_invalid_destination_name" {
  enabled = true
}

rule "aws_cloudwatch_log_group_invalid_kms_key_id" {
  enabled = true
}

rule "aws_cloudwatch_log_group_invalid_name" {
  enabled = true
}

rule "aws_cloudwatch_log_metric_filter_invalid_log_group_name" {
  enabled = true
}

rule "aws_cloudwatch_log_metric_filter_invalid_name" {
  enabled = true
}

rule "aws_cloudwatch_log_metric_filter_invalid_pattern" {
  enabled = true
}

rule "aws_cloudwatch_log_resource_policy_invalid_policy_document" {
  enabled = true
}

rule "aws_cloudwatch_log_stream_invalid_log_group_name" {
  enabled = true
}

rule "aws_cloudwatch_log_stream_invalid_name" {
  enabled = true
}

rule "aws_cloudwatch_log_subscription_filter_invalid_distribution" {
  enabled = true
}

rule "aws_cloudwatch_log_subscription_filter_invalid_filter_pattern" {
  enabled = true
}

rule "aws_cloudwatch_log_subscription_filter_invalid_log_group_name" {
  enabled = true
}

rule "aws_cloudwatch_log_subscription_filter_invalid_name" {
  enabled = true
}

rule "aws_cloudwatch_metric_alarm_invalid_alarm_description" {
  enabled = true
}

rule "aws_cloudwatch_metric_alarm_invalid_alarm_name" {
  enabled = true
}

rule "aws_cloudwatch_metric_alarm_invalid_comparison_operator" {
  enabled = true
}

rule "aws_cloudwatch_metric_alarm_invalid_evaluate_low_sample_count_percentiles" {
  enabled = true
}

rule "aws_cloudwatch_metric_alarm_invalid_metric_name" {
  enabled = true
}

rule "aws_cloudwatch_metric_alarm_invalid_namespace" {
  enabled = true
}

rule "aws_cloudwatch_metric_alarm_invalid_statistic" {
  enabled = true
}

rule "aws_cloudwatch_metric_alarm_invalid_treat_missing_data" {
  enabled = true
}

rule "aws_cloudwatch_metric_alarm_invalid_unit" {
  enabled = true
}

rule "aws_codeartifact_domain_invalid_domain" {
  enabled = true
}

rule "aws_codeartifact_domain_invalid_encryption_key" {
  enabled = true
}

rule "aws_codeartifact_domain_permissions_policy_invalid_domain" {
  enabled = true
}

rule "aws_codeartifact_domain_permissions_policy_invalid_domain_owner" {
  enabled = true
}

rule "aws_codeartifact_domain_permissions_policy_invalid_policy_document" {
  enabled = true
}

rule "aws_codeartifact_domain_permissions_policy_invalid_policy_revision" {
  enabled = true
}

rule "aws_codeartifact_repository_invalid_description" {
  enabled = true
}

rule "aws_codeartifact_repository_invalid_domain" {
  enabled = true
}

rule "aws_codeartifact_repository_invalid_domain_owner" {
  enabled = true
}

rule "aws_codeartifact_repository_invalid_repository" {
  enabled = true
}

rule "aws_codeartifact_repository_permissions_policy_invalid_domain" {
  enabled = true
}

rule "aws_codeartifact_repository_permissions_policy_invalid_domain_owner" {
  enabled = true
}

rule "aws_codeartifact_repository_permissions_policy_invalid_policy_document" {
  enabled = true
}

rule "aws_codeartifact_repository_permissions_policy_invalid_policy_revision" {
  enabled = true
}

rule "aws_codeartifact_repository_permissions_policy_invalid_repository" {
  enabled = true
}

rule "aws_codebuild_project_invalid_description" {
  enabled = true
}

rule "aws_codebuild_report_group_invalid_name" {
  enabled = true
}

rule "aws_codebuild_report_group_invalid_type" {
  enabled = true
}

rule "aws_codebuild_source_credential_invalid_auth_type" {
  enabled = true
}

rule "aws_codebuild_source_credential_invalid_server_type" {
  enabled = true
}

rule "aws_codecommit_approval_rule_template_association_invalid_approval_rule_template_name" {
  enabled = true
}

rule "aws_codecommit_approval_rule_template_association_invalid_repository_name" {
  enabled = true
}

rule "aws_codecommit_approval_rule_template_invalid_content" {
  enabled = true
}

rule "aws_codecommit_approval_rule_template_invalid_description" {
  enabled = true
}

rule "aws_codecommit_approval_rule_template_invalid_name" {
  enabled = true
}

rule "aws_codecommit_repository_invalid_default_branch" {
  enabled = true
}

rule "aws_codecommit_repository_invalid_description" {
  enabled = true
}

rule "aws_codecommit_repository_invalid_repository_name" {
  enabled = true
}

rule "aws_codecommit_trigger_invalid_repository_name" {
  enabled = true
}

rule "aws_codedeploy_app_invalid_compute_platform" {
  enabled = true
}

rule "aws_codedeploy_app_invalid_name" {
  enabled = true
}

rule "aws_codedeploy_deployment_config_invalid_compute_platform" {
  enabled = true
}

rule "aws_codedeploy_deployment_config_invalid_deployment_config_name" {
  enabled = true
}

rule "aws_codedeploy_deployment_group_invalid_app_name" {
  enabled = true
}

rule "aws_codedeploy_deployment_group_invalid_deployment_config_name" {
  enabled = true
}

rule "aws_codedeploy_deployment_group_invalid_deployment_group_name" {
  enabled = true
}

rule "aws_codepipeline_invalid_name" {
  enabled = true
}

rule "aws_codepipeline_invalid_role_arn" {
  enabled = true
}

rule "aws_codepipeline_webhook_invalid_authentication" {
  enabled = true
}

rule "aws_codepipeline_webhook_invalid_name" {
  enabled = true
}

rule "aws_codepipeline_webhook_invalid_target_action" {
  enabled = true
}

rule "aws_codepipeline_webhook_invalid_target_pipeline" {
  enabled = true
}

rule "aws_codestarconnections_connection_invalid_host_arn" {
  enabled = true
}

rule "aws_codestarconnections_connection_invalid_name" {
  enabled = true
}

rule "aws_codestarconnections_connection_invalid_provider_type" {
  enabled = true
}

rule "aws_codestarconnections_host_invalid_name" {
  enabled = true
}

rule "aws_codestarconnections_host_invalid_provider_endpoint" {
  enabled = true
}

rule "aws_codestarconnections_host_invalid_provider_type" {
  enabled = true
}

rule "aws_cognito_identity_pool_invalid_developer_provider_name" {
  enabled = true
}

rule "aws_cognito_identity_pool_invalid_identity_pool_name" {
  enabled = true
}

rule "aws_cognito_identity_pool_roles_attachment_invalid_identity_pool_id" {
  enabled = true
}

rule "aws_cognito_identity_provider_invalid_provider_name" {
  enabled = true
}

rule "aws_cognito_identity_provider_invalid_provider_type" {
  enabled = true
}

rule "aws_cognito_identity_provider_invalid_user_pool_id" {
  enabled = true
}

rule "aws_cognito_resource_server_invalid_identifier" {
  enabled = true
}

rule "aws_cognito_resource_server_invalid_name" {
  enabled = true
}

rule "aws_cognito_user_group_invalid_description" {
  enabled = true
}

rule "aws_cognito_user_group_invalid_name" {
  enabled = true
}

rule "aws_cognito_user_group_invalid_role_arn" {
  enabled = true
}

rule "aws_cognito_user_group_invalid_user_pool_id" {
  enabled = true
}

rule "aws_cognito_user_pool_client_invalid_default_redirect_uri" {
  enabled = true
}

rule "aws_cognito_user_pool_client_invalid_name" {
  enabled = true
}

rule "aws_cognito_user_pool_client_invalid_user_pool_id" {
  enabled = true
}

rule "aws_cognito_user_pool_domain_invalid_certificate_arn" {
  enabled = true
}

rule "aws_cognito_user_pool_domain_invalid_user_pool_id" {
  enabled = true
}

rule "aws_cognito_user_pool_invalid_email_verification_message" {
  enabled = true
}

rule "aws_cognito_user_pool_invalid_email_verification_subject" {
  enabled = true
}

rule "aws_cognito_user_pool_invalid_mfa_configuration" {
  enabled = true
}

rule "aws_cognito_user_pool_invalid_name" {
  enabled = true
}

rule "aws_cognito_user_pool_invalid_sms_authentication_message" {
  enabled = true
}

rule "aws_cognito_user_pool_invalid_sms_verification_message" {
  enabled = true
}

rule "aws_cognito_user_pool_ui_customization_invalid_client_id" {
  enabled = true
}

rule "aws_cognito_user_pool_ui_customization_invalid_user_pool_id" {
  enabled = true
}

rule "aws_config_aggregate_authorization_invalid_account_id" {
  enabled = true
}

rule "aws_config_aggregate_authorization_invalid_region" {
  enabled = true
}

rule "aws_config_config_rule_invalid_description" {
  enabled = true
}

rule "aws_config_config_rule_invalid_input_parameters" {
  enabled = true
}

rule "aws_config_config_rule_invalid_maximum_execution_frequency" {
  enabled = true
}

rule "aws_config_config_rule_invalid_name" {
  enabled = true
}

rule "aws_config_configuration_aggregator_invalid_name" {
  enabled = true
}

rule "aws_config_configuration_recorder_invalid_name" {
  enabled = true
}

rule "aws_config_configuration_recorder_status_invalid_name" {
  enabled = true
}

rule "aws_config_conformance_pack_invalid_delivery_s3_bucket" {
  enabled = true
}

rule "aws_config_conformance_pack_invalid_delivery_s3_key_prefix" {
  enabled = true
}

rule "aws_config_conformance_pack_invalid_name" {
  enabled = true
}

rule "aws_config_conformance_pack_invalid_template_body" {
  enabled = true
}

rule "aws_config_conformance_pack_invalid_template_s3_uri" {
  enabled = true
}

rule "aws_config_delivery_channel_invalid_name" {
  enabled = true
}

rule "aws_config_organization_conformance_pack_invalid_delivery_s3_bucket" {
  enabled = true
}

rule "aws_config_organization_conformance_pack_invalid_delivery_s3_key_prefix" {
  enabled = true
}

rule "aws_config_organization_conformance_pack_invalid_name" {
  enabled = true
}

rule "aws_config_organization_conformance_pack_invalid_template_body" {
  enabled = true
}

rule "aws_config_organization_conformance_pack_invalid_template_s3_uri" {
  enabled = true
}

rule "aws_config_organization_custom_rule_invalid_description" {
  enabled = true
}

rule "aws_config_organization_custom_rule_invalid_input_parameters" {
  enabled = true
}

rule "aws_config_organization_custom_rule_invalid_lambda_function_arn" {
  enabled = true
}

rule "aws_config_organization_custom_rule_invalid_maximum_execution_frequency" {
  enabled = true
}

rule "aws_config_organization_custom_rule_invalid_name" {
  enabled = true
}

rule "aws_config_organization_custom_rule_invalid_resource_id_scope" {
  enabled = true
}

rule "aws_config_organization_custom_rule_invalid_tag_key_scope" {
  enabled = true
}

rule "aws_config_organization_custom_rule_invalid_tag_value_scope" {
  enabled = true
}

rule "aws_config_organization_managed_rule_invalid_description" {
  enabled = true
}

rule "aws_config_organization_managed_rule_invalid_input_parameters" {
  enabled = true
}

rule "aws_config_organization_managed_rule_invalid_maximum_execution_frequency" {
  enabled = true
}

rule "aws_config_organization_managed_rule_invalid_name" {
  enabled = true
}

rule "aws_config_organization_managed_rule_invalid_resource_id_scope" {
  enabled = true
}

rule "aws_config_organization_managed_rule_invalid_rule_identifier" {
  enabled = true
}

rule "aws_config_organization_managed_rule_invalid_tag_key_scope" {
  enabled = true
}

rule "aws_config_organization_managed_rule_invalid_tag_value_scope" {
  enabled = true
}

rule "aws_config_remediation_configuration_invalid_config_rule_name" {
  enabled = true
}

rule "aws_config_remediation_configuration_invalid_target_id" {
  enabled = true
}

rule "aws_config_remediation_configuration_invalid_target_type" {
  enabled = true
}

rule "aws_connect_bot_association_invalid_instance_id" {
  enabled = true
}

rule "aws_connect_contact_flow_invalid_instance_id" {
  enabled = true
}

rule "aws_connect_contact_flow_invalid_type" {
  enabled = true
}

rule "aws_connect_hours_of_operation_invalid_description" {
  enabled = true
}

rule "aws_connect_hours_of_operation_invalid_instance_id" {
  enabled = true
}

rule "aws_connect_hours_of_operation_invalid_name" {
  enabled = true
}

rule "aws_connect_instance_invalid_identity_management_type" {
  enabled = true
}

rule "aws_connect_lambda_function_association_invalid_function_arn" {
  enabled = true
}

rule "aws_connect_lambda_function_association_invalid_instance_id" {
  enabled = true
}

rule "aws_cur_report_definition_invalid_compression" {
  enabled = true
}

rule "aws_cur_report_definition_invalid_format" {
  enabled = true
}

rule "aws_cur_report_definition_invalid_report_name" {
  enabled = true
}

rule "aws_cur_report_definition_invalid_s3_bucket" {
  enabled = true
}

rule "aws_cur_report_definition_invalid_s3_prefix" {
  enabled = true
}

rule "aws_cur_report_definition_invalid_s3_region" {
  enabled = true
}

rule "aws_cur_report_definition_invalid_time_unit" {
  enabled = true
}

rule "aws_customer_gateway_invalid_type" {
  enabled = true
}

rule "aws_datasync_agent_invalid_activation_key" {
  enabled = true
}

rule "aws_datasync_agent_invalid_name" {
  enabled = true
}

rule "aws_datasync_location_efs_invalid_efs_file_system_arn" {
  enabled = true
}

rule "aws_datasync_location_efs_invalid_subdirectory" {
  enabled = true
}

rule "aws_datasync_location_fsx_windows_file_system_invalid_domain" {
  enabled = true
}

rule "aws_datasync_location_fsx_windows_file_system_invalid_fsx_filesystem_arn" {
  enabled = true
}

rule "aws_datasync_location_fsx_windows_file_system_invalid_password" {
  enabled = true
}

rule "aws_datasync_location_fsx_windows_file_system_invalid_subdirectory" {
  enabled = true
}

rule "aws_datasync_location_fsx_windows_file_system_invalid_user" {
  enabled = true
}

rule "aws_datasync_location_nfs_invalid_server_hostname" {
  enabled = true
}

rule "aws_datasync_location_nfs_invalid_subdirectory" {
  enabled = true
}

rule "aws_datasync_location_s3_invalid_s3_bucket_arn" {
  enabled = true
}

rule "aws_datasync_location_s3_invalid_subdirectory" {
  enabled = true
}

rule "aws_datasync_location_smb_invalid_domain" {
  enabled = true
}

rule "aws_datasync_location_smb_invalid_password" {
  enabled = true
}

rule "aws_datasync_location_smb_invalid_server_hostname" {
  enabled = true
}

rule "aws_datasync_location_smb_invalid_subdirectory" {
  enabled = true
}

rule "aws_datasync_location_smb_invalid_user" {
  enabled = true
}

rule "aws_datasync_task_invalid_cloudwatch_log_group_arn" {
  enabled = true
}

rule "aws_datasync_task_invalid_destination_location_arn" {
  enabled = true
}

rule "aws_datasync_task_invalid_name" {
  enabled = true
}

rule "aws_datasync_task_invalid_source_location_arn" {
  enabled = true
}

rule "aws_db_proxy_endpoint_invalid_db_proxy_endpoint_name" {
  enabled = true
}

rule "aws_db_proxy_endpoint_invalid_db_proxy_name" {
  enabled = true
}

rule "aws_db_proxy_endpoint_invalid_target_role" {
  enabled = true
}

rule "aws_db_proxy_invalid_engine_family" {
  enabled = true
}

rule "aws_devicefarm_device_pool_invalid_description" {
  enabled = true
}

rule "aws_devicefarm_device_pool_invalid_name" {
  enabled = true
}

rule "aws_devicefarm_device_pool_invalid_project_arn" {
  enabled = true
}

rule "aws_devicefarm_network_profile_invalid_description" {
  enabled = true
}

rule "aws_devicefarm_network_profile_invalid_name" {
  enabled = true
}

rule "aws_devicefarm_network_profile_invalid_project_arn" {
  enabled = true
}

rule "aws_devicefarm_network_profile_invalid_type" {
  enabled = true
}

rule "aws_devicefarm_project_invalid_name" {
  enabled = true
}

rule "aws_devicefarm_upload_invalid_content_type" {
  enabled = true
}

rule "aws_devicefarm_upload_invalid_name" {
  enabled = true
}

rule "aws_devicefarm_upload_invalid_project_arn" {
  enabled = true
}

rule "aws_devicefarm_upload_invalid_type" {
  enabled = true
}

rule "aws_directory_service_conditional_forwarder_invalid_directory_id" {
  enabled = true
}

rule "aws_directory_service_conditional_forwarder_invalid_remote_domain_name" {
  enabled = true
}

rule "aws_directory_service_directory_invalid_description" {
  enabled = true
}

rule "aws_directory_service_directory_invalid_edition" {
  enabled = true
}

rule "aws_directory_service_directory_invalid_name" {
  enabled = true
}

rule "aws_directory_service_directory_invalid_password" {
  enabled = true
}

rule "aws_directory_service_directory_invalid_short_name" {
  enabled = true
}

rule "aws_directory_service_directory_invalid_size" {
  enabled = true
}

rule "aws_directory_service_directory_invalid_type" {
  enabled = true
}

rule "aws_directory_service_log_subscription_invalid_directory_id" {
  enabled = true
}

rule "aws_directory_service_log_subscription_invalid_log_group_name" {
  enabled = true
}

rule "aws_dlm_lifecycle_policy_invalid_description" {
  enabled = true
}

rule "aws_dlm_lifecycle_policy_invalid_execution_role_arn" {
  enabled = true
}

rule "aws_dlm_lifecycle_policy_invalid_state" {
  enabled = true
}

rule "aws_dms_endpoint_invalid_endpoint_type" {
  enabled = true
}

rule "aws_dms_endpoint_invalid_ssl_mode" {
  enabled = true
}

rule "aws_dms_replication_task_invalid_migration_type" {
  enabled = true
}

rule "aws_docdb_global_cluster_invalid_global_cluster_identifier" {
  enabled = true
}

rule "aws_dx_bgp_peer_invalid_address_family" {
  enabled = true
}

rule "aws_dx_hosted_private_virtual_interface_invalid_address_family" {
  enabled = true
}

rule "aws_dx_hosted_public_virtual_interface_invalid_address_family" {
  enabled = true
}

rule "aws_dx_hosted_transit_virtual_interface_invalid_address_family" {
  enabled = true
}

rule "aws_dx_private_virtual_interface_invalid_address_family" {
  enabled = true
}

rule "aws_dx_public_virtual_interface_invalid_address_family" {
  enabled = true
}

rule "aws_dynamodb_global_table_invalid_name" {
  enabled = true
}

rule "aws_dynamodb_kinesis_streaming_destination_invalid_stream_arn" {
  enabled = true
}

rule "aws_dynamodb_kinesis_streaming_destination_invalid_table_name" {
  enabled = true
}

rule "aws_dynamodb_table_invalid_billing_mode" {
  enabled = true
}

rule "aws_dynamodb_table_invalid_hash_key" {
  enabled = true
}

rule "aws_dynamodb_table_invalid_name" {
  enabled = true
}

rule "aws_dynamodb_table_invalid_range_key" {
  enabled = true
}

rule "aws_dynamodb_table_item_invalid_hash_key" {
  enabled = true
}

rule "aws_dynamodb_table_item_invalid_range_key" {
  enabled = true
}

rule "aws_dynamodb_table_item_invalid_table_name" {
  enabled = true
}

rule "aws_dynamodb_tag_invalid_resource_arn" {
  enabled = true
}

rule "aws_ebs_volume_invalid_type" {
  enabled = true
}

rule "aws_ec2_availability_zone_group_invalid_opt_in_status" {
  enabled = true
}

rule "aws_ec2_capacity_reservation_invalid_end_date_type" {
  enabled = true
}

rule "aws_ec2_capacity_reservation_invalid_instance_match_criteria" {
  enabled = true
}

rule "aws_ec2_capacity_reservation_invalid_instance_platform" {
  enabled = true
}

rule "aws_ec2_capacity_reservation_invalid_tenancy" {
  enabled = true
}

rule "aws_ec2_client_vpn_endpoint_invalid_transport_protocol" {
  enabled = true
}

rule "aws_ec2_fleet_invalid_excess_capacity_termination_policy" {
  enabled = true
}

rule "aws_ec2_fleet_invalid_type" {
  enabled = true
}

rule "aws_ec2_host_invalid_auto_placement" {
  enabled = true
}

rule "aws_ec2_host_invalid_host_recovery" {
  enabled = true
}

rule "aws_ec2_subnet_cidr_reservation_invalid_reservation_type" {
  enabled = true
}

rule "aws_ec2_traffic_mirror_filter_rule_invalid_rule_action" {
  enabled = true
}

rule "aws_ec2_traffic_mirror_filter_rule_invalid_traffic_direction" {
  enabled = true
}

rule "aws_ec2_transit_gateway_invalid_auto_accept_shared_attachments" {
  enabled = true
}

rule "aws_ec2_transit_gateway_invalid_default_route_table_association" {
  enabled = true
}

rule "aws_ec2_transit_gateway_invalid_default_route_table_propagation" {
  enabled = true
}

rule "aws_ec2_transit_gateway_invalid_dns_support" {
  enabled = true
}

rule "aws_ec2_transit_gateway_vpc_attachment_invalid_dns_support" {
  enabled = true
}

rule "aws_ec2_transit_gateway_vpc_attachment_invalid_ipv6_support" {
  enabled = true
}

rule "aws_ecr_lifecycle_policy_invalid_policy" {
  enabled = true
}

rule "aws_ecr_lifecycle_policy_invalid_repository" {
  enabled = true
}

rule "aws_ecr_pull_through_cache_rule_invalid_ecr_repository_prefix" {
  enabled = true
}

rule "aws_ecr_registry_policy_invalid_policy" {
  enabled = true
}

rule "aws_ecr_registry_scanning_configuration_invalid_scan_type" {
  enabled = true
}

rule "aws_ecr_repository_invalid_name" {
  enabled = true
}

rule "aws_ecr_repository_policy_invalid_policy" {
  enabled = true
}

rule "aws_ecr_repository_policy_invalid_repository" {
  enabled = true
}

rule "aws_ecrpublic_repository_invalid_repository_name" {
  enabled = true
}

rule "aws_ecrpublic_repository_policy_invalid_policy" {
  enabled = true
}

rule "aws_ecrpublic_repository_policy_invalid_repository_name" {
  enabled = true
}

rule "aws_ecs_account_setting_default_invalid_name" {
  enabled = true
}

rule "aws_ecs_service_invalid_launch_type" {
  enabled = true
}

rule "aws_ecs_service_invalid_propagate_tags" {
  enabled = true
}

rule "aws_ecs_service_invalid_scheduling_strategy" {
  enabled = true
}

rule "aws_ecs_task_definition_invalid_ipc_mode" {
  enabled = true
}

rule "aws_ecs_task_definition_invalid_network_mode" {
  enabled = true
}

rule "aws_ecs_task_definition_invalid_pid_mode" {
  enabled = true
}

rule "aws_ecs_task_set_invalid_launch_type" {
  enabled = true
}

rule "aws_efs_access_point_invalid_file_system_id" {
  enabled = true
}

rule "aws_efs_backup_policy_invalid_file_system_id" {
  enabled = true
}

rule "aws_efs_file_system_invalid_creation_token" {
  enabled = true
}

rule "aws_efs_file_system_invalid_kms_key_id" {
  enabled = true
}

rule "aws_efs_file_system_invalid_performance_mode" {
  enabled = true
}

rule "aws_efs_file_system_invalid_throughput_mode" {
  enabled = true
}

rule "aws_efs_file_system_policy_invalid_file_system_id" {
  enabled = true
}

rule "aws_efs_file_system_policy_invalid_policy" {
  enabled = true
}

rule "aws_efs_mount_target_invalid_file_system_id" {
  enabled = true
}

rule "aws_efs_mount_target_invalid_ip_address" {
  enabled = true
}

rule "aws_efs_mount_target_invalid_subnet_id" {
  enabled = true
}

rule "aws_eks_addon_invalid_cluster_name" {
  enabled = true
}

rule "aws_eks_addon_invalid_resolve_conflicts" {
  enabled = true
}

rule "aws_eks_addon_invalid_service_account_role_arn" {
  enabled = true
}

rule "aws_eks_cluster_invalid_name" {
  enabled = true
}

rule "aws_eks_node_group_invalid_ami_type" {
  enabled = true
}

rule "aws_eks_node_group_invalid_capacity_type" {
  enabled = true
}

rule "aws_elastic_beanstalk_application_invalid_description" {
  enabled = true
}

rule "aws_elastic_beanstalk_application_invalid_name" {
  enabled = true
}

rule "aws_elastic_beanstalk_application_version_invalid_application" {
  enabled = true
}

rule "aws_elastic_beanstalk_application_version_invalid_bucket" {
  enabled = true
}

rule "aws_elastic_beanstalk_application_version_invalid_description" {
  enabled = true
}

rule "aws_elastic_beanstalk_application_version_invalid_key" {
  enabled = true
}

rule "aws_elastic_beanstalk_application_version_invalid_name" {
  enabled = true
}

rule "aws_elastic_beanstalk_configuration_template_invalid_application" {
  enabled = true
}

rule "aws_elastic_beanstalk_configuration_template_invalid_description" {
  enabled = true
}

rule "aws_elastic_beanstalk_configuration_template_invalid_name" {
  enabled = true
}

rule "aws_elastic_beanstalk_environment_invalid_application" {
  enabled = true
}

rule "aws_elastic_beanstalk_environment_invalid_cname_prefix" {
  enabled = true
}

rule "aws_elastic_beanstalk_environment_invalid_description" {
  enabled = true
}

rule "aws_elastic_beanstalk_environment_invalid_name" {
  enabled = true
}

rule "aws_elastic_beanstalk_environment_invalid_template_name" {
  enabled = true
}

rule "aws_elastic_beanstalk_environment_invalid_version_label" {
  enabled = true
}

rule "aws_elasticache_cluster_invalid_az_mode" {
  enabled = true
}

rule "aws_elasticache_user_group_invalid_engine" {
  enabled = true
}

rule "aws_elasticache_user_invalid_access_string" {
  enabled = true
}

rule "aws_elasticache_user_invalid_engine" {
  enabled = true
}

rule "aws_elasticache_user_invalid_user_id" {
  enabled = true
}

rule "aws_elasticsearch_domain_invalid_domain_name" {
  enabled = true
}

rule "aws_elasticsearch_domain_invalid_elasticsearch_version" {
  enabled = true
}

rule "aws_elasticsearch_domain_policy_invalid_domain_name" {
  enabled = true
}

rule "aws_elasticsearch_domain_saml_options_invalid_domain_name" {
  enabled = true
}

rule "aws_elastictranscoder_pipeline_invalid_aws_kms_key_arn" {
  enabled = true
}

rule "aws_elastictranscoder_pipeline_invalid_input_bucket" {
  enabled = true
}

rule "aws_elastictranscoder_pipeline_invalid_name" {
  enabled = true
}

rule "aws_elastictranscoder_pipeline_invalid_output_bucket" {
  enabled = true
}

rule "aws_elastictranscoder_pipeline_invalid_role" {
  enabled = true
}

rule "aws_elastictranscoder_preset_invalid_container" {
  enabled = true
}

rule "aws_elastictranscoder_preset_invalid_description" {
  enabled = true
}

rule "aws_elastictranscoder_preset_invalid_name" {
  enabled = true
}

rule "aws_emr_cluster_invalid_scale_down_behavior" {
  enabled = true
}

rule "aws_emr_studio_invalid_auth_mode" {
  enabled = true
}

rule "aws_emr_studio_session_mapping_invalid_identity_type" {
  enabled = true
}

rule "aws_flow_log_invalid_log_destination_type" {
  enabled = true
}

rule "aws_flow_log_invalid_traffic_type" {
  enabled = true
}

rule "aws_fms_admin_account_invalid_account_id" {
  enabled = true
}

rule "aws_fms_policy_invalid_name" {
  enabled = true
}

rule "aws_fms_policy_invalid_resource_type" {
  enabled = true
}

rule "aws_fsx_backup_invalid_file_system_id" {
  enabled = true
}

rule "aws_fsx_backup_invalid_volume_id" {
  enabled = true
}

rule "aws_fsx_lustre_file_system_invalid_weekly_maintenance_start_time" {
  enabled = true
}

rule "aws_fsx_ontap_file_system_invalid_daily_automatic_backup_start_time" {
  enabled = true
}

rule "aws_fsx_ontap_file_system_invalid_deployment_type" {
  enabled = true
}

rule "aws_fsx_ontap_file_system_invalid_endpoint_ip_address_range" {
  enabled = true
}

rule "aws_fsx_ontap_file_system_invalid_fsx_admin_password" {
  enabled = true
}

rule "aws_fsx_ontap_file_system_invalid_preferred_subnet_id" {
  enabled = true
}

rule "aws_fsx_ontap_file_system_invalid_weekly_maintenance_start_time" {
  enabled = true
}

rule "aws_fsx_ontap_storage_virtual_machine_invalid_file_system_id" {
  enabled = true
}

rule "aws_fsx_ontap_storage_virtual_machine_invalid_name" {
  enabled = true
}

rule "aws_fsx_ontap_storage_virtual_machine_invalid_root_volume_security_style" {
  enabled = true
}

rule "aws_fsx_ontap_volume_invalid_junction_path" {
  enabled = true
}

rule "aws_fsx_ontap_volume_invalid_name" {
  enabled = true
}

rule "aws_fsx_ontap_volume_invalid_security_style" {
  enabled = true
}

rule "aws_fsx_ontap_volume_invalid_storage_virtual_machine_id" {
  enabled = true
}

rule "aws_fsx_openzfs_file_system_invalid_backup_id" {
  enabled = true
}

rule "aws_fsx_openzfs_file_system_invalid_daily_automatic_backup_start_time" {
  enabled = true
}

rule "aws_fsx_openzfs_file_system_invalid_deployment_type" {
  enabled = true
}

rule "aws_fsx_openzfs_file_system_invalid_storage_type" {
  enabled = true
}

rule "aws_fsx_openzfs_file_system_invalid_weekly_maintenance_start_time" {
  enabled = true
}

rule "aws_fsx_openzfs_snapshot_invalid_name" {
  enabled = true
}

rule "aws_fsx_openzfs_snapshot_invalid_volume_id" {
  enabled = true
}

rule "aws_fsx_openzfs_volume_invalid_data_compression_type" {
  enabled = true
}

rule "aws_fsx_openzfs_volume_invalid_parent_volume_id" {
  enabled = true
}

rule "aws_fsx_windows_file_system_invalid_active_directory_id" {
  enabled = true
}

rule "aws_fsx_windows_file_system_invalid_daily_automatic_backup_start_time" {
  enabled = true
}

rule "aws_fsx_windows_file_system_invalid_weekly_maintenance_start_time" {
  enabled = true
}

rule "aws_gamelift_alias_invalid_description" {
  enabled = true
}

rule "aws_gamelift_alias_invalid_name" {
  enabled = true
}

rule "aws_gamelift_build_invalid_name" {
  enabled = true
}

rule "aws_gamelift_build_invalid_operating_system" {
  enabled = true
}

rule "aws_gamelift_build_invalid_version" {
  enabled = true
}

rule "aws_gamelift_fleet_invalid_build_id" {
  enabled = true
}

rule "aws_gamelift_fleet_invalid_description" {
  enabled = true
}

rule "aws_gamelift_fleet_invalid_ec2_instance_type" {
  enabled = true
}

rule "aws_gamelift_fleet_invalid_name" {
  enabled = true
}

rule "aws_gamelift_fleet_invalid_new_game_session_protection_policy" {
  enabled = true
}

rule "aws_gamelift_game_session_queue_invalid_name" {
  enabled = true
}

rule "aws_globalaccelerator_accelerator_invalid_ip_address_type" {
  enabled = true
}

rule "aws_globalaccelerator_accelerator_invalid_name" {
  enabled = true
}

rule "aws_globalaccelerator_endpoint_group_invalid_health_check_path" {
  enabled = true
}

rule "aws_globalaccelerator_endpoint_group_invalid_health_check_protocol" {
  enabled = true
}

rule "aws_globalaccelerator_endpoint_group_invalid_listener_arn" {
  enabled = true
}

rule "aws_globalaccelerator_listener_invalid_accelerator_arn" {
  enabled = true
}

rule "aws_globalaccelerator_listener_invalid_client_affinity" {
  enabled = true
}

rule "aws_globalaccelerator_listener_invalid_protocol" {
  enabled = true
}

rule "aws_glue_catalog_table_invalid_table_type" {
  enabled = true
}

rule "aws_glue_catalog_table_invalid_view_expanded_text" {
  enabled = true
}

rule "aws_glue_catalog_table_invalid_view_original_text" {
  enabled = true
}

rule "aws_glue_connection_invalid_connection_type" {
  enabled = true
}

rule "aws_glue_crawler_invalid_security_configuration" {
  enabled = true
}

rule "aws_glue_crawler_invalid_table_prefix" {
  enabled = true
}

rule "aws_glue_dev_endpoint_invalid_role_arn" {
  enabled = true
}

rule "aws_glue_dev_endpoint_invalid_worker_type" {
  enabled = true
}

rule "aws_glue_ml_transform_invalid_glue_version" {
  enabled = true
}

rule "aws_glue_ml_transform_invalid_worker_type" {
  enabled = true
}

rule "aws_glue_registry_invalid_registry_name" {
  enabled = true
}

rule "aws_glue_resource_policy_invalid_enable_hybrid" {
  enabled = true
}

rule "aws_glue_schema_invalid_compatibility" {
  enabled = true
}

rule "aws_glue_schema_invalid_data_format" {
  enabled = true
}

rule "aws_glue_schema_invalid_schema_definition" {
  enabled = true
}

rule "aws_glue_schema_invalid_schema_name" {
  enabled = true
}

rule "aws_glue_trigger_invalid_type" {
  enabled = true
}

rule "aws_glue_user_defined_function_invalid_owner_type" {
  enabled = true
}

rule "aws_guardduty_detector_invalid_finding_publishing_frequency" {
  enabled = true
}

rule "aws_guardduty_filter_invalid_action" {
  enabled = true
}

rule "aws_guardduty_filter_invalid_description" {
  enabled = true
}

rule "aws_guardduty_filter_invalid_detector_id" {
  enabled = true
}

rule "aws_guardduty_filter_invalid_name" {
  enabled = true
}

rule "aws_guardduty_invite_accepter_invalid_detector_id" {
  enabled = true
}

rule "aws_guardduty_ipset_invalid_detector_id" {
  enabled = true
}

rule "aws_guardduty_ipset_invalid_format" {
  enabled = true
}

rule "aws_guardduty_ipset_invalid_location" {
  enabled = true
}

rule "aws_guardduty_ipset_invalid_name" {
  enabled = true
}

rule "aws_guardduty_member_invalid_detector_id" {
  enabled = true
}

rule "aws_guardduty_member_invalid_email" {
  enabled = true
}

rule "aws_guardduty_organization_configuration_invalid_detector_id" {
  enabled = true
}

rule "aws_guardduty_publishing_destination_invalid_destination_type" {
  enabled = true
}

rule "aws_guardduty_publishing_destination_invalid_detector_id" {
  enabled = true
}

rule "aws_guardduty_threatintelset_invalid_detector_id" {
  enabled = true
}

rule "aws_guardduty_threatintelset_invalid_format" {
  enabled = true
}

rule "aws_guardduty_threatintelset_invalid_location" {
  enabled = true
}

rule "aws_guardduty_threatintelset_invalid_name" {
  enabled = true
}

rule "aws_iam_access_key_invalid_status" {
  enabled = true
}

rule "aws_iam_access_key_invalid_user" {
  enabled = true
}

rule "aws_iam_group_invalid_name" {
  enabled = true
}

rule "aws_iam_group_invalid_path" {
  enabled = true
}

rule "aws_iam_group_membership_invalid_group" {
  enabled = true
}

rule "aws_iam_group_policy_attachment_invalid_group" {
  enabled = true
}

rule "aws_iam_group_policy_attachment_invalid_policy_arn" {
  enabled = true
}

rule "aws_iam_group_policy_invalid_group" {
  enabled = true
}

rule "aws_iam_group_policy_invalid_name" {
  enabled = true
}

rule "aws_iam_group_policy_invalid_policy" {
  enabled = true
}

rule "aws_iam_instance_profile_invalid_name" {
  enabled = true
}

rule "aws_iam_instance_profile_invalid_path" {
  enabled = true
}

rule "aws_iam_instance_profile_invalid_role" {
  enabled = true
}

rule "aws_iam_openid_connect_provider_invalid_url" {
  enabled = true
}

rule "aws_iam_policy_attachment_invalid_policy_arn" {
  enabled = true
}

rule "aws_iam_policy_invalid_description" {
  enabled = true
}

rule "aws_iam_policy_invalid_name" {
  enabled = true
}

rule "aws_iam_policy_invalid_path" {
  enabled = true
}

rule "aws_iam_policy_invalid_policy" {
  enabled = true
}

rule "aws_iam_role_invalid_assume_role_policy" {
  enabled = true
}

rule "aws_iam_role_invalid_description" {
  enabled = true
}

rule "aws_iam_role_invalid_name" {
  enabled = true
}

rule "aws_iam_role_invalid_path" {
  enabled = true
}

rule "aws_iam_role_invalid_permissions_boundary" {
  enabled = true
}

rule "aws_iam_role_policy_attachment_invalid_policy_arn" {
  enabled = true
}

rule "aws_iam_role_policy_attachment_invalid_role" {
  enabled = true
}

rule "aws_iam_role_policy_invalid_name" {
  enabled = true
}

rule "aws_iam_role_policy_invalid_policy" {
  enabled = true
}

rule "aws_iam_role_policy_invalid_role" {
  enabled = true
}

rule "aws_iam_saml_provider_invalid_name" {
  enabled = true
}

rule "aws_iam_saml_provider_invalid_saml_metadata_document" {
  enabled = true
}

rule "aws_iam_server_certificate_invalid_certificate_body" {
  enabled = true
}

rule "aws_iam_server_certificate_invalid_certificate_chain" {
  enabled = true
}

rule "aws_iam_server_certificate_invalid_name" {
  enabled = true
}

rule "aws_iam_server_certificate_invalid_path" {
  enabled = true
}

rule "aws_iam_server_certificate_invalid_private_key" {
  enabled = true
}

rule "aws_iam_service_linked_role_invalid_aws_service_name" {
  enabled = true
}

rule "aws_iam_service_linked_role_invalid_custom_suffix" {
  enabled = true
}

rule "aws_iam_service_linked_role_invalid_description" {
  enabled = true
}

rule "aws_iam_user_group_membership_invalid_user" {
  enabled = true
}

rule "aws_iam_user_invalid_name" {
  enabled = true
}

rule "aws_iam_user_invalid_path" {
  enabled = true
}

rule "aws_iam_user_invalid_permissions_boundary" {
  enabled = true
}

rule "aws_iam_user_login_profile_invalid_user" {
  enabled = true
}

rule "aws_iam_user_policy_attachment_invalid_policy_arn" {
  enabled = true
}

rule "aws_iam_user_policy_attachment_invalid_user" {
  enabled = true
}

rule "aws_iam_user_policy_invalid_name" {
  enabled = true
}

rule "aws_iam_user_policy_invalid_policy" {
  enabled = true
}

rule "aws_iam_user_policy_invalid_user" {
  enabled = true
}

rule "aws_iam_user_ssh_key_invalid_encoding" {
  enabled = true
}

rule "aws_iam_user_ssh_key_invalid_public_key" {
  enabled = true
}

rule "aws_iam_user_ssh_key_invalid_status" {
  enabled = true
}

rule "aws_iam_user_ssh_key_invalid_username" {
  enabled = true
}

rule "aws_imagebuilder_component_invalid_change_description" {
  enabled = true
}

rule "aws_imagebuilder_component_invalid_data" {
  enabled = true
}

rule "aws_imagebuilder_component_invalid_description" {
  enabled = true
}

rule "aws_imagebuilder_component_invalid_kms_key_id" {
  enabled = true
}

rule "aws_imagebuilder_component_invalid_name" {
  enabled = true
}

rule "aws_imagebuilder_component_invalid_platform" {
  enabled = true
}

rule "aws_imagebuilder_component_invalid_version" {
  enabled = true
}

rule "aws_imagebuilder_distribution_configuration_invalid_description" {
  enabled = true
}

rule "aws_imagebuilder_distribution_configuration_invalid_name" {
  enabled = true
}

rule "aws_imagebuilder_image_invalid_distribution_configuration_arn" {
  enabled = true
}

rule "aws_imagebuilder_image_invalid_image_recipe_arn" {
  enabled = true
}

rule "aws_imagebuilder_image_invalid_infrastructure_configuration_arn" {
  enabled = true
}

rule "aws_imagebuilder_image_pipeline_invalid_description" {
  enabled = true
}

rule "aws_imagebuilder_image_pipeline_invalid_distribution_configuration_arn" {
  enabled = true
}

rule "aws_imagebuilder_image_pipeline_invalid_image_recipe_arn" {
  enabled = true
}

rule "aws_imagebuilder_image_pipeline_invalid_infrastructure_configuration_arn" {
  enabled = true
}

rule "aws_imagebuilder_image_pipeline_invalid_name" {
  enabled = true
}

rule "aws_imagebuilder_image_pipeline_invalid_status" {
  enabled = true
}

rule "aws_imagebuilder_image_recipe_invalid_description" {
  enabled = true
}

rule "aws_imagebuilder_image_recipe_invalid_name" {
  enabled = true
}

rule "aws_imagebuilder_image_recipe_invalid_parent_image" {
  enabled = true
}

rule "aws_imagebuilder_image_recipe_invalid_version" {
  enabled = true
}

rule "aws_imagebuilder_image_recipe_invalid_working_directory" {
  enabled = true
}

rule "aws_imagebuilder_infrastructure_configuration_invalid_description" {
  enabled = true
}

rule "aws_imagebuilder_infrastructure_configuration_invalid_instance_profile_name" {
  enabled = true
}

rule "aws_imagebuilder_infrastructure_configuration_invalid_key_pair" {
  enabled = true
}

rule "aws_imagebuilder_infrastructure_configuration_invalid_name" {
  enabled = true
}

rule "aws_imagebuilder_infrastructure_configuration_invalid_sns_topic_arn" {
  enabled = true
}

rule "aws_imagebuilder_infrastructure_configuration_invalid_subnet_id" {
  enabled = true
}

rule "aws_inspector_assessment_target_invalid_name" {
  enabled = true
}

rule "aws_inspector_assessment_target_invalid_resource_group_arn" {
  enabled = true
}

rule "aws_inspector_assessment_template_invalid_name" {
  enabled = true
}

rule "aws_inspector_assessment_template_invalid_target_arn" {
  enabled = true
}

rule "aws_instance_invalid_instance_initiated_shutdown_behavior" {
  enabled = true
}

rule "aws_instance_invalid_tenancy" {
  enabled = true
}

rule "aws_instance_invalid_type" {
  enabled = true
}

rule "aws_iot_certificate_invalid_csr" {
  enabled = true
}

rule "aws_iot_policy_attachment_invalid_policy" {
  enabled = true
}

rule "aws_iot_policy_invalid_name" {
  enabled = true
}

rule "aws_iot_policy_invalid_policy" {
  enabled = true
}

rule "aws_iot_role_alias_invalid_alias" {
  enabled = true
}

rule "aws_iot_role_alias_invalid_role_arn" {
  enabled = true
}

rule "aws_iot_thing_invalid_name" {
  enabled = true
}

rule "aws_iot_thing_invalid_thing_type_name" {
  enabled = true
}

rule "aws_iot_thing_principal_attachment_invalid_thing" {
  enabled = true
}

rule "aws_iot_thing_type_invalid_name" {
  enabled = true
}

rule "aws_iot_topic_rule_invalid_name" {
  enabled = true
}

rule "aws_kinesis_analytics_application_invalid_code" {
  enabled = true
}

rule "aws_kinesis_analytics_application_invalid_description" {
  enabled = true
}

rule "aws_kinesis_analytics_application_invalid_name" {
  enabled = true
}

rule "aws_kinesis_firehose_delivery_stream_invalid_name" {
  enabled = true
}

rule "aws_kinesis_stream_invalid_encryption_type" {
  enabled = true
}

rule "aws_kinesis_stream_invalid_kms_key_id" {
  enabled = true
}

rule "aws_kinesis_stream_invalid_name" {
  enabled = true
}

rule "aws_kinesisanalyticsv2_application_invalid_description" {
  enabled = true
}

rule "aws_kinesisanalyticsv2_application_invalid_name" {
  enabled = true
}

rule "aws_kinesisanalyticsv2_application_invalid_runtime_environment" {
  enabled = true
}

rule "aws_kinesisanalyticsv2_application_invalid_service_execution_role" {
  enabled = true
}

rule "aws_kinesisanalyticsv2_application_snapshot_invalid_application_name" {
  enabled = true
}

rule "aws_kinesisanalyticsv2_application_snapshot_invalid_snapshot_name" {
  enabled = true
}

rule "aws_kms_alias_invalid_name" {
  enabled = true
}

rule "aws_kms_alias_invalid_target_key_id" {
  enabled = true
}

rule "aws_kms_ciphertext_invalid_key_id" {
  enabled = true
}

rule "aws_kms_external_key_invalid_description" {
  enabled = true
}

rule "aws_kms_external_key_invalid_policy" {
  enabled = true
}

rule "aws_kms_grant_invalid_grantee_principal" {
  enabled = true
}

rule "aws_kms_grant_invalid_key_id" {
  enabled = true
}

rule "aws_kms_grant_invalid_name" {
  enabled = true
}

rule "aws_kms_grant_invalid_retiring_principal" {
  enabled = true
}

rule "aws_kms_key_invalid_description" {
  enabled = true
}

rule "aws_kms_key_invalid_key_usage" {
  enabled = true
}

rule "aws_kms_key_invalid_policy" {
  enabled = true
}

rule "aws_kms_replica_external_key_invalid_description" {
  enabled = true
}

rule "aws_kms_replica_external_key_invalid_policy" {
  enabled = true
}

rule "aws_kms_replica_key_invalid_description" {
  enabled = true
}

rule "aws_kms_replica_key_invalid_policy" {
  enabled = true
}

rule "aws_lakeformation_resource_invalid_role_arn" {
  enabled = true
}

rule "aws_lambda_alias_invalid_description" {
  enabled = true
}

rule "aws_lambda_alias_invalid_function_name" {
  enabled = true
}

rule "aws_lambda_alias_invalid_function_version" {
  enabled = true
}

rule "aws_lambda_code_signing_config_invalid_description" {
  enabled = true
}

rule "aws_lambda_event_source_mapping_invalid_event_source_arn" {
  enabled = true
}

rule "aws_lambda_event_source_mapping_invalid_function_name" {
  enabled = true
}

rule "aws_lambda_event_source_mapping_invalid_starting_position" {
  enabled = true
}

rule "aws_lambda_function_event_invoke_config_invalid_function_name" {
  enabled = true
}

rule "aws_lambda_function_event_invoke_config_invalid_qualifier" {
  enabled = true
}

rule "aws_lambda_function_invalid_description" {
  enabled = true
}

rule "aws_lambda_function_invalid_function_name" {
  enabled = true
}

rule "aws_lambda_function_invalid_handler" {
  enabled = true
}

rule "aws_lambda_function_invalid_kms_key_arn" {
  enabled = true
}

rule "aws_lambda_function_invalid_role" {
  enabled = true
}

rule "aws_lambda_function_invalid_runtime" {
  enabled = true
}

rule "aws_lambda_function_invalid_s3_key" {
  enabled = true
}

rule "aws_lambda_function_invalid_s3_object_version" {
  enabled = true
}

rule "aws_lambda_layer_version_invalid_description" {
  enabled = true
}

rule "aws_lambda_layer_version_invalid_layer_name" {
  enabled = true
}

rule "aws_lambda_layer_version_invalid_license_info" {
  enabled = true
}

rule "aws_lambda_layer_version_invalid_s3_key" {
  enabled = true
}

rule "aws_lambda_layer_version_invalid_s3_object_version" {
  enabled = true
}

rule "aws_lambda_layer_version_permission_invalid_action" {
  enabled = true
}

rule "aws_lambda_layer_version_permission_invalid_layer_name" {
  enabled = true
}

rule "aws_lambda_layer_version_permission_invalid_organization_id" {
  enabled = true
}

rule "aws_lambda_layer_version_permission_invalid_principal" {
  enabled = true
}

rule "aws_lambda_layer_version_permission_invalid_statement_id" {
  enabled = true
}

rule "aws_lambda_permission_invalid_action" {
  enabled = true
}

rule "aws_lambda_permission_invalid_event_source_token" {
  enabled = true
}

rule "aws_lambda_permission_invalid_function_name" {
  enabled = true
}

rule "aws_lambda_permission_invalid_principal" {
  enabled = true
}

rule "aws_lambda_permission_invalid_qualifier" {
  enabled = true
}

rule "aws_lambda_permission_invalid_source_account" {
  enabled = true
}

rule "aws_lambda_permission_invalid_source_arn" {
  enabled = true
}

rule "aws_lambda_permission_invalid_statement_id" {
  enabled = true
}

rule "aws_lambda_provisioned_concurrency_config_invalid_function_name" {
  enabled = true
}

rule "aws_lambda_provisioned_concurrency_config_invalid_qualifier" {
  enabled = true
}

rule "aws_launch_configuration_invalid_spot_price" {
  enabled = true
}

rule "aws_launch_configuration_invalid_type" {
  enabled = true
}

rule "aws_launch_template_invalid_description" {
  enabled = true
}

rule "aws_launch_template_invalid_instance_initiated_shutdown_behavior" {
  enabled = true
}

rule "aws_launch_template_invalid_instance_type" {
  enabled = true
}

rule "aws_launch_template_invalid_name" {
  enabled = true
}

rule "aws_lb_invalid_ip_address_type" {
  enabled = true
}

rule "aws_lb_invalid_load_balancer_type" {
  enabled = true
}

rule "aws_lb_listener_invalid_protocol" {
  enabled = true
}

rule "aws_lb_target_group_invalid_protocol" {
  enabled = true
}

rule "aws_lb_target_group_invalid_target_type" {
  enabled = true
}

rule "aws_licensemanager_license_configuration_invalid_license_counting_type" {
  enabled = true
}

rule "aws_lightsail_instance_invalid_blueprint_id" {
  enabled = true
}

rule "aws_lightsail_instance_invalid_bundle_id" {
  enabled = true
}

rule "aws_lightsail_instance_invalid_key_pair_name" {
  enabled = true
}

rule "aws_lightsail_instance_public_ports_invalid_instance_name" {
  enabled = true
}

rule "aws_lightsail_key_pair_invalid_name" {
  enabled = true
}

rule "aws_lightsail_static_ip_attachment_invalid_instance_name" {
  enabled = true
}

rule "aws_lightsail_static_ip_attachment_invalid_static_ip_name" {
  enabled = true
}

rule "aws_lightsail_static_ip_invalid_name" {
  enabled = true
}

rule "aws_macie2_account_invalid_finding_publishing_frequency" {
  enabled = true
}

rule "aws_macie2_account_invalid_status" {
  enabled = true
}

rule "aws_macie2_classification_job_invalid_job_status" {
  enabled = true
}

rule "aws_macie2_classification_job_invalid_job_type" {
  enabled = true
}

rule "aws_macie2_findings_filter_invalid_action" {
  enabled = true
}

rule "aws_macie_member_account_association_invalid_member_account_id" {
  enabled = true
}

rule "aws_macie_s3_bucket_association_invalid_bucket_name" {
  enabled = true
}

rule "aws_macie_s3_bucket_association_invalid_member_account_id" {
  enabled = true
}

rule "aws_macie_s3_bucket_association_invalid_prefix" {
  enabled = true
}

rule "aws_media_store_container_invalid_name" {
  enabled = true
}

rule "aws_media_store_container_policy_invalid_container_name" {
  enabled = true
}

rule "aws_memorydb_cluster_invalid_acl_name" {
  enabled = true
}

rule "aws_memorydb_user_invalid_access_string" {
  enabled = true
}

rule "aws_memorydb_user_invalid_user_name" {
  enabled = true
}

rule "aws_mq_broker_invalid_deployment_mode" {
  enabled = true
}

rule "aws_msk_cluster_invalid_cluster_name" {
  enabled = true
}

rule "aws_msk_cluster_invalid_enhanced_monitoring" {
  enabled = true
}

rule "aws_msk_cluster_invalid_kafka_version" {
  enabled = true
}

rule "aws_network_acl_rule_invalid_rule_action" {
  enabled = true
}

rule "aws_networkfirewall_firewall_invalid_description" {
  enabled = true
}

rule "aws_networkfirewall_firewall_invalid_firewall_policy_arn" {
  enabled = true
}

rule "aws_networkfirewall_firewall_invalid_name" {
  enabled = true
}

rule "aws_networkfirewall_firewall_invalid_vpc_id" {
  enabled = true
}

rule "aws_networkfirewall_firewall_policy_invalid_description" {
  enabled = true
}

rule "aws_networkfirewall_firewall_policy_invalid_name" {
  enabled = true
}

rule "aws_networkfirewall_logging_configuration_invalid_firewall_arn" {
  enabled = true
}

rule "aws_networkfirewall_resource_policy_invalid_policy" {
  enabled = true
}

rule "aws_networkfirewall_resource_policy_invalid_resource_arn" {
  enabled = true
}

rule "aws_networkfirewall_rule_group_invalid_description" {
  enabled = true
}

rule "aws_networkfirewall_rule_group_invalid_name" {
  enabled = true
}

rule "aws_networkfirewall_rule_group_invalid_rules" {
  enabled = true
}

rule "aws_networkfirewall_rule_group_invalid_type" {
  enabled = true
}

rule "aws_opsworks_application_invalid_type" {
  enabled = true
}

rule "aws_opsworks_instance_invalid_architecture" {
  enabled = true
}

rule "aws_opsworks_instance_invalid_auto_scaling_type" {
  enabled = true
}

rule "aws_opsworks_instance_invalid_root_device_type" {
  enabled = true
}

rule "aws_opsworks_stack_invalid_default_root_device_type" {
  enabled = true
}

rule "aws_organizations_account_invalid_email" {
  enabled = true
}

rule "aws_organizations_account_invalid_iam_user_access_to_billing" {
  enabled = true
}

rule "aws_organizations_account_invalid_name" {
  enabled = true
}

rule "aws_organizations_account_invalid_parent_id" {
  enabled = true
}

rule "aws_organizations_account_invalid_role_name" {
  enabled = true
}

rule "aws_organizations_delegated_administrator_invalid_account_id" {
  enabled = true
}

rule "aws_organizations_delegated_administrator_invalid_service_principal" {
  enabled = true
}

rule "aws_organizations_organization_invalid_feature_set" {
  enabled = true
}

rule "aws_organizations_organizational_unit_invalid_name" {
  enabled = true
}

rule "aws_organizations_organizational_unit_invalid_parent_id" {
  enabled = true
}

rule "aws_organizations_policy_attachment_invalid_policy_id" {
  enabled = true
}

rule "aws_organizations_policy_attachment_invalid_target_id" {
  enabled = true
}

rule "aws_organizations_policy_invalid_content" {
  enabled = true
}

rule "aws_organizations_policy_invalid_description" {
  enabled = true
}

rule "aws_organizations_policy_invalid_name" {
  enabled = true
}

rule "aws_organizations_policy_invalid_type" {
  enabled = true
}

rule "aws_placement_group_invalid_strategy" {
  enabled = true
}

rule "aws_prometheus_alert_manager_definition_invalid_workspace_id" {
  enabled = true
}

rule "aws_prometheus_rule_group_namespace_invalid_name" {
  enabled = true
}

rule "aws_prometheus_rule_group_namespace_invalid_workspace_id" {
  enabled = true
}

rule "aws_prometheus_workspace_invalid_alias" {
  enabled = true
}

rule "aws_quicksight_data_source_invalid_aws_account_id" {
  enabled = true
}

rule "aws_quicksight_data_source_invalid_name" {
  enabled = true
}

rule "aws_quicksight_data_source_invalid_type" {
  enabled = true
}

rule "aws_quicksight_group_invalid_aws_account_id" {
  enabled = true
}

rule "aws_quicksight_group_invalid_description" {
  enabled = true
}

rule "aws_quicksight_group_invalid_group_name" {
  enabled = true
}

rule "aws_quicksight_group_invalid_namespace" {
  enabled = true
}

rule "aws_quicksight_group_membership_invalid_aws_account_id" {
  enabled = true
}

rule "aws_quicksight_group_membership_invalid_group_name" {
  enabled = true
}

rule "aws_quicksight_group_membership_invalid_member_name" {
  enabled = true
}

rule "aws_quicksight_group_membership_invalid_namespace" {
  enabled = true
}

rule "aws_quicksight_user_invalid_aws_account_id" {
  enabled = true
}

rule "aws_quicksight_user_invalid_identity_type" {
  enabled = true
}

rule "aws_quicksight_user_invalid_namespace" {
  enabled = true
}

rule "aws_quicksight_user_invalid_session_name" {
  enabled = true
}

rule "aws_quicksight_user_invalid_user_name" {
  enabled = true
}

rule "aws_quicksight_user_invalid_user_role" {
  enabled = true
}

rule "aws_rds_cluster_role_association_invalid_db_cluster_identifier" {
  enabled = true
}

rule "aws_redshift_cluster_invalid_availability_zone" {
  enabled = true
}

rule "aws_redshift_cluster_invalid_cluster_identifier" {
  enabled = true
}

rule "aws_redshift_cluster_invalid_cluster_parameter_group_name" {
  enabled = true
}

rule "aws_redshift_cluster_invalid_cluster_subnet_group_name" {
  enabled = true
}

rule "aws_redshift_cluster_invalid_cluster_type" {
  enabled = true
}

rule "aws_redshift_cluster_invalid_cluster_version" {
  enabled = true
}

rule "aws_redshift_cluster_invalid_database_name" {
  enabled = true
}

rule "aws_redshift_cluster_invalid_elastic_ip" {
  enabled = true
}

rule "aws_redshift_cluster_invalid_final_snapshot_identifier" {
  enabled = true
}

rule "aws_redshift_cluster_invalid_kms_key_id" {
  enabled = true
}

rule "aws_redshift_cluster_invalid_master_password" {
  enabled = true
}

rule "aws_redshift_cluster_invalid_master_username" {
  enabled = true
}

rule "aws_redshift_cluster_invalid_node_type" {
  enabled = true
}

rule "aws_redshift_cluster_invalid_owner_account" {
  enabled = true
}

rule "aws_redshift_cluster_invalid_preferred_maintenance_window" {
  enabled = true
}

rule "aws_redshift_cluster_invalid_snapshot_cluster_identifier" {
  enabled = true
}

rule "aws_redshift_cluster_invalid_snapshot_identifier" {
  enabled = true
}

rule "aws_redshift_event_subscription_invalid_name" {
  enabled = true
}

rule "aws_redshift_event_subscription_invalid_severity" {
  enabled = true
}

rule "aws_redshift_event_subscription_invalid_sns_topic_arn" {
  enabled = true
}

rule "aws_redshift_event_subscription_invalid_source_type" {
  enabled = true
}

rule "aws_redshift_parameter_group_invalid_description" {
  enabled = true
}

rule "aws_redshift_parameter_group_invalid_family" {
  enabled = true
}

rule "aws_redshift_parameter_group_invalid_name" {
  enabled = true
}

rule "aws_redshift_security_group_invalid_description" {
  enabled = true
}

rule "aws_redshift_security_group_invalid_name" {
  enabled = true
}

rule "aws_redshift_snapshot_copy_grant_invalid_kms_key_id" {
  enabled = true
}

rule "aws_redshift_snapshot_copy_grant_invalid_snapshot_copy_grant_name" {
  enabled = true
}

rule "aws_redshift_snapshot_schedule_association_invalid_cluster_identifier" {
  enabled = true
}

rule "aws_redshift_snapshot_schedule_association_invalid_schedule_identifier" {
  enabled = true
}

rule "aws_redshift_snapshot_schedule_invalid_description" {
  enabled = true
}

rule "aws_redshift_snapshot_schedule_invalid_identifier" {
  enabled = true
}

rule "aws_redshift_snapshot_schedule_invalid_identifier_prefix" {
  enabled = true
}

rule "aws_redshift_subnet_group_invalid_description" {
  enabled = true
}

rule "aws_redshift_subnet_group_invalid_name" {
  enabled = true
}

rule "aws_resourcegroups_group_invalid_name" {
  enabled = true
}

rule "aws_route53_delegation_set_invalid_reference_name" {
  enabled = true
}

rule "aws_route53_health_check_invalid_cloudwatch_alarm_name" {
  enabled = true
}

rule "aws_route53_health_check_invalid_cloudwatch_alarm_region" {
  enabled = true
}

rule "aws_route53_health_check_invalid_fqdn" {
  enabled = true
}

rule "aws_route53_health_check_invalid_insufficient_data_health_status" {
  enabled = true
}

rule "aws_route53_health_check_invalid_ip_address" {
  enabled = true
}

rule "aws_route53_health_check_invalid_reference_name" {
  enabled = true
}

rule "aws_route53_health_check_invalid_resource_path" {
  enabled = true
}

rule "aws_route53_health_check_invalid_search_string" {
  enabled = true
}

rule "aws_route53_health_check_invalid_type" {
  enabled = true
}

rule "aws_route53_query_log_invalid_zone_id" {
  enabled = true
}

rule "aws_route53_record_invalid_health_check_id" {
  enabled = true
}

rule "aws_route53_record_invalid_name" {
  enabled = true
}

rule "aws_route53_record_invalid_set_identifier" {
  enabled = true
}

rule "aws_route53_record_invalid_type" {
  enabled = true
}

rule "aws_route53_record_invalid_zone_id" {
  enabled = true
}

rule "aws_route53_resolver_dnssec_config_invalid_resource_id" {
  enabled = true
}

rule "aws_route53_resolver_endpoint_invalid_direction" {
  enabled = true
}

rule "aws_route53_resolver_firewall_config_invalid_firewall_fail_open" {
  enabled = true
}

rule "aws_route53_resolver_firewall_config_invalid_resource_id" {
  enabled = true
}

rule "aws_route53_resolver_firewall_rule_group_association_invalid_firewall_rule_group_id" {
  enabled = true
}

rule "aws_route53_resolver_firewall_rule_group_association_invalid_mutation_protection" {
  enabled = true
}

rule "aws_route53_resolver_firewall_rule_group_association_invalid_vpc_id" {
  enabled = true
}

rule "aws_route53_resolver_firewall_rule_invalid_action" {
  enabled = true
}

rule "aws_route53_resolver_firewall_rule_invalid_block_override_dns_type" {
  enabled = true
}

rule "aws_route53_resolver_firewall_rule_invalid_block_override_domain" {
  enabled = true
}

rule "aws_route53_resolver_firewall_rule_invalid_block_response" {
  enabled = true
}

rule "aws_route53_resolver_firewall_rule_invalid_firewall_domain_list_id" {
  enabled = true
}

rule "aws_route53_resolver_firewall_rule_invalid_firewall_rule_group_id" {
  enabled = true
}

rule "aws_route53_resolver_query_log_config_association_invalid_resolver_query_log_config_id" {
  enabled = true
}

rule "aws_route53_resolver_query_log_config_association_invalid_resource_id" {
  enabled = true
}

rule "aws_route53_resolver_query_log_config_invalid_destination_arn" {
  enabled = true
}

rule "aws_route53_resolver_rule_association_invalid_resolver_rule_id" {
  enabled = true
}

rule "aws_route53_resolver_rule_association_invalid_vpc_id" {
  enabled = true
}

rule "aws_route53_resolver_rule_invalid_domain_name" {
  enabled = true
}

rule "aws_route53_resolver_rule_invalid_resolver_endpoint_id" {
  enabled = true
}

rule "aws_route53_resolver_rule_invalid_rule_type" {
  enabled = true
}

rule "aws_route53_zone_association_invalid_vpc_id" {
  enabled = true
}

rule "aws_route53_zone_association_invalid_vpc_region" {
  enabled = true
}

rule "aws_route53_zone_association_invalid_zone_id" {
  enabled = true
}

rule "aws_route53_zone_invalid_comment" {
  enabled = true
}

rule "aws_route53_zone_invalid_delegation_set_id" {
  enabled = true
}

rule "aws_route53_zone_invalid_name" {
  enabled = true
}

rule "aws_route53recoverycontrolconfig_cluster_invalid_name" {
  enabled = true
}

rule "aws_route53recoverycontrolconfig_control_panel_invalid_cluster_arn" {
  enabled = true
}

rule "aws_route53recoverycontrolconfig_control_panel_invalid_name" {
  enabled = true
}

rule "aws_route53recoverycontrolconfig_routing_control_invalid_cluster_arn" {
  enabled = true
}

rule "aws_route53recoverycontrolconfig_routing_control_invalid_control_panel_arn" {
  enabled = true
}

rule "aws_route53recoverycontrolconfig_routing_control_invalid_name" {
  enabled = true
}

rule "aws_route53recoverycontrolconfig_safety_rule_invalid_control_panel_arn" {
  enabled = true
}

rule "aws_route53recoverycontrolconfig_safety_rule_invalid_name" {
  enabled = true
}

rule "aws_route53recoveryreadiness_resource_set_invalid_resource_set_type" {
  enabled = true
}

rule "aws_s3_bucket_intelligent_tiering_configuration_invalid_status" {
  enabled = true
}

rule "aws_s3_bucket_invalid_acceleration_status" {
  enabled = true
}

rule "aws_s3_bucket_invalid_request_payer" {
  enabled = true
}

rule "aws_s3_bucket_inventory_invalid_included_object_versions" {
  enabled = true
}

rule "aws_s3_bucket_object_invalid_acl" {
  enabled = true
}

rule "aws_s3_bucket_object_invalid_server_side_encryption" {
  enabled = true
}

rule "aws_s3_bucket_object_invalid_storage_class" {
  enabled = true
}

rule "aws_s3_object_copy_invalid_acl" {
  enabled = true
}

rule "aws_s3_object_copy_invalid_metadata_directive" {
  enabled = true
}

rule "aws_s3_object_copy_invalid_object_lock_legal_hold_status" {
  enabled = true
}

rule "aws_s3_object_copy_invalid_object_lock_mode" {
  enabled = true
}

rule "aws_s3_object_copy_invalid_request_payer" {
  enabled = true
}

rule "aws_s3_object_copy_invalid_server_side_encryption" {
  enabled = true
}

rule "aws_s3_object_copy_invalid_source" {
  enabled = true
}

rule "aws_s3_object_copy_invalid_storage_class" {
  enabled = true
}

rule "aws_s3_object_copy_invalid_tagging_directive" {
  enabled = true
}

rule "aws_s3control_access_point_policy_invalid_access_point_arn" {
  enabled = true
}

rule "aws_s3control_bucket_invalid_bucket" {
  enabled = true
}

rule "aws_s3control_bucket_invalid_outpost_id" {
  enabled = true
}

rule "aws_s3control_bucket_lifecycle_configuration_invalid_bucket" {
  enabled = true
}

rule "aws_s3control_bucket_policy_invalid_bucket" {
  enabled = true
}

rule "aws_s3control_multi_region_access_point_invalid_account_id" {
  enabled = true
}

rule "aws_s3control_multi_region_access_point_policy_invalid_account_id" {
  enabled = true
}

rule "aws_s3control_object_lambda_access_point_invalid_account_id" {
  enabled = true
}

rule "aws_s3control_object_lambda_access_point_invalid_name" {
  enabled = true
}

rule "aws_s3control_object_lambda_access_point_policy_invalid_account_id" {
  enabled = true
}

rule "aws_s3control_object_lambda_access_point_policy_invalid_name" {
  enabled = true
}

rule "aws_sagemaker_app_invalid_app_name" {
  enabled = true
}

rule "aws_sagemaker_app_invalid_app_type" {
  enabled = true
}

rule "aws_sagemaker_app_invalid_domain_id" {
  enabled = true
}

rule "aws_sagemaker_app_invalid_user_profile_name" {
  enabled = true
}

rule "aws_sagemaker_code_repository_invalid_code_repository_name" {
  enabled = true
}

rule "aws_sagemaker_device_fleet_invalid_description" {
  enabled = true
}

rule "aws_sagemaker_device_fleet_invalid_device_fleet_name" {
  enabled = true
}

rule "aws_sagemaker_device_fleet_invalid_role_arn" {
  enabled = true
}

rule "aws_sagemaker_domain_invalid_app_network_access_type" {
  enabled = true
}

rule "aws_sagemaker_domain_invalid_auth_mode" {
  enabled = true
}

rule "aws_sagemaker_domain_invalid_domain_name" {
  enabled = true
}

rule "aws_sagemaker_domain_invalid_kms_key_id" {
  enabled = true
}

rule "aws_sagemaker_domain_invalid_vpc_id" {
  enabled = true
}

rule "aws_sagemaker_endpoint_configuration_invalid_kms_key_arn" {
  enabled = true
}

rule "aws_sagemaker_endpoint_configuration_invalid_name" {
  enabled = true
}

rule "aws_sagemaker_endpoint_invalid_endpoint_config_name" {
  enabled = true
}

rule "aws_sagemaker_endpoint_invalid_name" {
  enabled = true
}

rule "aws_sagemaker_feature_group_invalid_description" {
  enabled = true
}

rule "aws_sagemaker_feature_group_invalid_event_time_feature_name" {
  enabled = true
}

rule "aws_sagemaker_feature_group_invalid_feature_group_name" {
  enabled = true
}

rule "aws_sagemaker_feature_group_invalid_record_identifier_feature_name" {
  enabled = true
}

rule "aws_sagemaker_feature_group_invalid_role_arn" {
  enabled = true
}

rule "aws_sagemaker_flow_definition_invalid_flow_definition_name" {
  enabled = true
}

rule "aws_sagemaker_flow_definition_invalid_role_arn" {
  enabled = true
}

rule "aws_sagemaker_human_task_ui_invalid_human_task_ui_name" {
  enabled = true
}

rule "aws_sagemaker_image_invalid_description" {
  enabled = true
}

rule "aws_sagemaker_image_invalid_display_name" {
  enabled = true
}

rule "aws_sagemaker_image_invalid_image_name" {
  enabled = true
}

rule "aws_sagemaker_image_invalid_role_arn" {
  enabled = true
}

rule "aws_sagemaker_image_version_invalid_base_image" {
  enabled = true
}

rule "aws_sagemaker_image_version_invalid_image_name" {
  enabled = true
}

rule "aws_sagemaker_model_invalid_execution_role_arn" {
  enabled = true
}

rule "aws_sagemaker_model_invalid_name" {
  enabled = true
}

rule "aws_sagemaker_model_package_group_invalid_model_package_group_description" {
  enabled = true
}

rule "aws_sagemaker_model_package_group_invalid_model_package_group_name" {
  enabled = true
}

rule "aws_sagemaker_model_package_group_policy_invalid_model_package_group_name" {
  enabled = true
}

rule "aws_sagemaker_notebook_instance_invalid_instance_type" {
  enabled = true
}

rule "aws_sagemaker_notebook_instance_invalid_kms_key_id" {
  enabled = true
}

rule "aws_sagemaker_notebook_instance_invalid_lifecycle_config_name" {
  enabled = true
}

rule "aws_sagemaker_notebook_instance_invalid_name" {
  enabled = true
}

rule "aws_sagemaker_notebook_instance_invalid_role_arn" {
  enabled = true
}

rule "aws_sagemaker_notebook_instance_invalid_subnet_id" {
  enabled = true
}

rule "aws_sagemaker_notebook_instance_lifecycle_configuration_invalid_name" {
  enabled = true
}

rule "aws_sagemaker_studio_lifecycle_config_invalid_studio_lifecycle_config_app_type" {
  enabled = true
}

rule "aws_sagemaker_studio_lifecycle_config_invalid_studio_lifecycle_config_content" {
  enabled = true
}

rule "aws_sagemaker_studio_lifecycle_config_invalid_studio_lifecycle_config_name" {
  enabled = true
}

rule "aws_sagemaker_user_profile_invalid_domain_id" {
  enabled = true
}

rule "aws_sagemaker_user_profile_invalid_single_sign_on_user_identifier" {
  enabled = true
}

rule "aws_sagemaker_user_profile_invalid_single_sign_on_user_value" {
  enabled = true
}

rule "aws_sagemaker_user_profile_invalid_user_profile_name" {
  enabled = true
}

rule "aws_sagemaker_workforce_invalid_workforce_name" {
  enabled = true
}

rule "aws_sagemaker_workteam_invalid_description" {
  enabled = true
}

rule "aws_sagemaker_workteam_invalid_workforce_name" {
  enabled = true
}

rule "aws_sagemaker_workteam_invalid_workteam_name" {
  enabled = true
}

rule "aws_schemas_discoverer_invalid_description" {
  enabled = true
}

rule "aws_schemas_discoverer_invalid_source_arn" {
  enabled = true
}

rule "aws_schemas_registry_invalid_description" {
  enabled = true
}

rule "aws_schemas_schema_invalid_content" {
  enabled = true
}

rule "aws_schemas_schema_invalid_description" {
  enabled = true
}

rule "aws_schemas_schema_invalid_type" {
  enabled = true
}

rule "aws_secretsmanager_secret_invalid_description" {
  enabled = true
}

rule "aws_secretsmanager_secret_invalid_kms_key_id" {
  enabled = true
}

rule "aws_secretsmanager_secret_invalid_name" {
  enabled = true
}

rule "aws_secretsmanager_secret_invalid_policy" {
  enabled = true
}

rule "aws_secretsmanager_secret_invalid_rotation_lambda_arn" {
  enabled = true
}

rule "aws_secretsmanager_secret_policy_invalid_policy" {
  enabled = true
}

rule "aws_secretsmanager_secret_policy_invalid_secret_arn" {
  enabled = true
}

rule "aws_secretsmanager_secret_rotation_invalid_rotation_lambda_arn" {
  enabled = true
}

rule "aws_secretsmanager_secret_rotation_invalid_secret_id" {
  enabled = true
}

rule "aws_secretsmanager_secret_version_invalid_secret_id" {
  enabled = true
}

rule "aws_secretsmanager_secret_version_invalid_secret_string" {
  enabled = true
}

rule "aws_securityhub_action_target_invalid_description" {
  enabled = true
}

rule "aws_securityhub_action_target_invalid_identifier" {
  enabled = true
}

rule "aws_securityhub_action_target_invalid_name" {
  enabled = true
}

rule "aws_securityhub_finding_aggregator_invalid_linking_mode" {
  enabled = true
}

rule "aws_securityhub_insight_invalid_group_by_attribute" {
  enabled = true
}

rule "aws_securityhub_insight_invalid_name" {
  enabled = true
}

rule "aws_securityhub_invite_accepter_invalid_master_id" {
  enabled = true
}

rule "aws_securityhub_member_invalid_email" {
  enabled = true
}

rule "aws_securityhub_organization_admin_account_invalid_admin_account_id" {
  enabled = true
}

rule "aws_securityhub_product_subscription_invalid_product_arn" {
  enabled = true
}

rule "aws_securityhub_standards_control_invalid_control_status" {
  enabled = true
}

rule "aws_securityhub_standards_control_invalid_disabled_reason" {
  enabled = true
}

rule "aws_securityhub_standards_control_invalid_standards_control_arn" {
  enabled = true
}

rule "aws_securityhub_standards_subscription_invalid_standards_arn" {
  enabled = true
}

rule "aws_service_discovery_http_namespace_invalid_description" {
  enabled = true
}

rule "aws_service_discovery_http_namespace_invalid_name" {
  enabled = true
}

rule "aws_service_discovery_instance_invalid_instance_id" {
  enabled = true
}

rule "aws_service_discovery_instance_invalid_service_id" {
  enabled = true
}

rule "aws_service_discovery_private_dns_namespace_invalid_description" {
  enabled = true
}

rule "aws_service_discovery_private_dns_namespace_invalid_name" {
  enabled = true
}

rule "aws_service_discovery_private_dns_namespace_invalid_vpc" {
  enabled = true
}

rule "aws_service_discovery_public_dns_namespace_invalid_description" {
  enabled = true
}

rule "aws_service_discovery_public_dns_namespace_invalid_name" {
  enabled = true
}

rule "aws_service_discovery_service_invalid_description" {
  enabled = true
}

rule "aws_servicecatalog_budget_resource_association_invalid_budget_name" {
  enabled = true
}

rule "aws_servicecatalog_budget_resource_association_invalid_resource_id" {
  enabled = true
}

rule "aws_servicecatalog_constraint_invalid_accept_language" {
  enabled = true
}

rule "aws_servicecatalog_constraint_invalid_description" {
  enabled = true
}

rule "aws_servicecatalog_constraint_invalid_portfolio_id" {
  enabled = true
}

rule "aws_servicecatalog_constraint_invalid_product_id" {
  enabled = true
}

rule "aws_servicecatalog_constraint_invalid_type" {
  enabled = true
}

rule "aws_servicecatalog_portfolio_invalid_description" {
  enabled = true
}

rule "aws_servicecatalog_portfolio_invalid_name" {
  enabled = true
}

rule "aws_servicecatalog_portfolio_invalid_provider_name" {
  enabled = true
}

rule "aws_servicecatalog_portfolio_share_invalid_accept_language" {
  enabled = true
}

rule "aws_servicecatalog_portfolio_share_invalid_portfolio_id" {
  enabled = true
}

rule "aws_servicecatalog_portfolio_share_invalid_principal_id" {
  enabled = true
}

rule "aws_servicecatalog_portfolio_share_invalid_type" {
  enabled = true
}

rule "aws_servicecatalog_principal_portfolio_association_invalid_accept_language" {
  enabled = true
}

rule "aws_servicecatalog_principal_portfolio_association_invalid_portfolio_id" {
  enabled = true
}

rule "aws_servicecatalog_principal_portfolio_association_invalid_principal_arn" {
  enabled = true
}

rule "aws_servicecatalog_principal_portfolio_association_invalid_principal_type" {
  enabled = true
}

rule "aws_servicecatalog_product_invalid_accept_language" {
  enabled = true
}

rule "aws_servicecatalog_product_invalid_description" {
  enabled = true
}

rule "aws_servicecatalog_product_invalid_distributor" {
  enabled = true
}

rule "aws_servicecatalog_product_invalid_name" {
  enabled = true
}

rule "aws_servicecatalog_product_invalid_owner" {
  enabled = true
}

rule "aws_servicecatalog_product_invalid_support_description" {
  enabled = true
}

rule "aws_servicecatalog_product_invalid_support_email" {
  enabled = true
}

rule "aws_servicecatalog_product_invalid_support_url" {
  enabled = true
}

rule "aws_servicecatalog_product_invalid_type" {
  enabled = true
}

rule "aws_servicecatalog_product_portfolio_association_invalid_accept_language" {
  enabled = true
}

rule "aws_servicecatalog_product_portfolio_association_invalid_portfolio_id" {
  enabled = true
}

rule "aws_servicecatalog_product_portfolio_association_invalid_product_id" {
  enabled = true
}

rule "aws_servicecatalog_product_portfolio_association_invalid_source_portfolio_id" {
  enabled = true
}

rule "aws_servicecatalog_provisioned_product_invalid_accept_language" {
  enabled = true
}

rule "aws_servicecatalog_provisioned_product_invalid_name" {
  enabled = true
}

rule "aws_servicecatalog_provisioned_product_invalid_path_id" {
  enabled = true
}

rule "aws_servicecatalog_provisioned_product_invalid_path_name" {
  enabled = true
}

rule "aws_servicecatalog_provisioned_product_invalid_product_id" {
  enabled = true
}

rule "aws_servicecatalog_provisioned_product_invalid_product_name" {
  enabled = true
}

rule "aws_servicecatalog_provisioned_product_invalid_provisioning_artifact_id" {
  enabled = true
}

rule "aws_servicecatalog_provisioned_product_invalid_provisioning_artifact_name" {
  enabled = true
}

rule "aws_servicecatalog_provisioning_artifact_invalid_accept_language" {
  enabled = true
}

rule "aws_servicecatalog_provisioning_artifact_invalid_description" {
  enabled = true
}

rule "aws_servicecatalog_provisioning_artifact_invalid_guidance" {
  enabled = true
}

rule "aws_servicecatalog_provisioning_artifact_invalid_name" {
  enabled = true
}

rule "aws_servicecatalog_provisioning_artifact_invalid_product_id" {
  enabled = true
}

rule "aws_servicecatalog_provisioning_artifact_invalid_type" {
  enabled = true
}

rule "aws_servicecatalog_service_action_invalid_accept_language" {
  enabled = true
}

rule "aws_servicecatalog_service_action_invalid_description" {
  enabled = true
}

rule "aws_servicecatalog_service_action_invalid_name" {
  enabled = true
}

rule "aws_servicecatalog_tag_option_invalid_key" {
  enabled = true
}

rule "aws_servicecatalog_tag_option_invalid_value" {
  enabled = true
}

rule "aws_servicecatalog_tag_option_resource_association_invalid_tag_option_id" {
  enabled = true
}

rule "aws_servicequotas_service_quota_invalid_quota_code" {
  enabled = true
}

rule "aws_servicequotas_service_quota_invalid_service_code" {
  enabled = true
}

rule "aws_ses_domain_mail_from_invalid_behavior_on_mx_failure" {
  enabled = true
}

rule "aws_ses_identity_notification_topic_invalid_notification_type" {
  enabled = true
}

rule "aws_ses_identity_policy_invalid_name" {
  enabled = true
}

rule "aws_ses_receipt_filter_invalid_policy" {
  enabled = true
}

rule "aws_ses_receipt_rule_invalid_tls_policy" {
  enabled = true
}

rule "aws_sfn_activity_invalid_name" {
  enabled = true
}

rule "aws_sfn_state_machine_invalid_definition" {
  enabled = true
}

rule "aws_sfn_state_machine_invalid_name" {
  enabled = true
}

rule "aws_sfn_state_machine_invalid_role_arn" {
  enabled = true
}

rule "aws_shield_protection_group_invalid_aggregation" {
  enabled = true
}

rule "aws_shield_protection_group_invalid_pattern" {
  enabled = true
}

rule "aws_shield_protection_group_invalid_protection_group_id" {
  enabled = true
}

rule "aws_shield_protection_group_invalid_resource_type" {
  enabled = true
}

rule "aws_shield_protection_invalid_name" {
  enabled = true
}

rule "aws_shield_protection_invalid_resource_arn" {
  enabled = true
}

rule "aws_signer_signing_job_invalid_profile_name" {
  enabled = true
}

rule "aws_signer_signing_profile_invalid_name" {
  enabled = true
}

rule "aws_signer_signing_profile_permission_invalid_profile_name" {
  enabled = true
}

rule "aws_signer_signing_profile_permission_invalid_profile_version" {
  enabled = true
}

rule "aws_spot_fleet_request_invalid_allocation_strategy" {
  enabled = true
}

rule "aws_spot_fleet_request_invalid_fleet_type" {
  enabled = true
}

rule "aws_spot_fleet_request_invalid_instance_interruption_behaviour" {
  enabled = true
}

rule "aws_spot_instance_request_invalid_instance_interruption_behavior" {
  enabled = true
}

rule "aws_ssm_activation_invalid_description" {
  enabled = true
}

rule "aws_ssm_activation_invalid_iam_role" {
  enabled = true
}

rule "aws_ssm_activation_invalid_name" {
  enabled = true
}

rule "aws_ssm_association_invalid_association_name" {
  enabled = true
}

rule "aws_ssm_association_invalid_compliance_severity" {
  enabled = true
}

rule "aws_ssm_association_invalid_document_version" {
  enabled = true
}

rule "aws_ssm_association_invalid_instance_id" {
  enabled = true
}

rule "aws_ssm_association_invalid_max_concurrency" {
  enabled = true
}

rule "aws_ssm_association_invalid_max_errors" {
  enabled = true
}

rule "aws_ssm_association_invalid_name" {
  enabled = true
}

rule "aws_ssm_association_invalid_schedule_expression" {
  enabled = true
}

rule "aws_ssm_document_invalid_document_format" {
  enabled = true
}

rule "aws_ssm_document_invalid_document_type" {
  enabled = true
}

rule "aws_ssm_document_invalid_name" {
  enabled = true
}

rule "aws_ssm_maintenance_window_invalid_name" {
  enabled = true
}

rule "aws_ssm_maintenance_window_invalid_schedule" {
  enabled = true
}

rule "aws_ssm_maintenance_window_target_invalid_description" {
  enabled = true
}

rule "aws_ssm_maintenance_window_target_invalid_name" {
  enabled = true
}

rule "aws_ssm_maintenance_window_target_invalid_owner_information" {
  enabled = true
}

rule "aws_ssm_maintenance_window_target_invalid_resource_type" {
  enabled = true
}

rule "aws_ssm_maintenance_window_target_invalid_window_id" {
  enabled = true
}

rule "aws_ssm_maintenance_window_task_invalid_description" {
  enabled = true
}

rule "aws_ssm_maintenance_window_task_invalid_max_concurrency" {
  enabled = true
}

rule "aws_ssm_maintenance_window_task_invalid_max_errors" {
  enabled = true
}

rule "aws_ssm_maintenance_window_task_invalid_name" {
  enabled = true
}

rule "aws_ssm_maintenance_window_task_invalid_task_arn" {
  enabled = true
}

rule "aws_ssm_maintenance_window_task_invalid_task_type" {
  enabled = true
}

rule "aws_ssm_maintenance_window_task_invalid_window_id" {
  enabled = true
}

rule "aws_ssm_parameter_invalid_allowed_pattern" {
  enabled = true
}

rule "aws_ssm_parameter_invalid_description" {
  enabled = true
}

rule "aws_ssm_parameter_invalid_key_id" {
  enabled = true
}

rule "aws_ssm_parameter_invalid_name" {
  enabled = true
}

rule "aws_ssm_parameter_invalid_tier" {
  enabled = true
}

rule "aws_ssm_parameter_invalid_type" {
  enabled = true
}

rule "aws_ssm_patch_baseline_invalid_approved_patches_compliance_level" {
  enabled = true
}

rule "aws_ssm_patch_baseline_invalid_description" {
  enabled = true
}

rule "aws_ssm_patch_baseline_invalid_name" {
  enabled = true
}

rule "aws_ssm_patch_baseline_invalid_operating_system" {
  enabled = true
}

rule "aws_ssm_patch_group_invalid_baseline_id" {
  enabled = true
}

rule "aws_ssm_patch_group_invalid_patch_group" {
  enabled = true
}

rule "aws_ssm_resource_data_sync_invalid_name" {
  enabled = true
}

rule "aws_ssoadmin_account_assignment_invalid_instance_arn" {
  enabled = true
}

rule "aws_ssoadmin_account_assignment_invalid_permission_set_arn" {
  enabled = true
}

rule "aws_ssoadmin_account_assignment_invalid_principal_id" {
  enabled = true
}

rule "aws_ssoadmin_account_assignment_invalid_principal_type" {
  enabled = true
}

rule "aws_ssoadmin_account_assignment_invalid_target_id" {
  enabled = true
}

rule "aws_ssoadmin_account_assignment_invalid_target_type" {
  enabled = true
}

rule "aws_ssoadmin_managed_policy_attachment_invalid_instance_arn" {
  enabled = true
}

rule "aws_ssoadmin_managed_policy_attachment_invalid_managed_policy_arn" {
  enabled = true
}

rule "aws_ssoadmin_managed_policy_attachment_invalid_permission_set_arn" {
  enabled = true
}

rule "aws_ssoadmin_permission_set_inline_policy_invalid_inline_policy" {
  enabled = true
}

rule "aws_ssoadmin_permission_set_inline_policy_invalid_instance_arn" {
  enabled = true
}

rule "aws_ssoadmin_permission_set_inline_policy_invalid_permission_set_arn" {
  enabled = true
}

rule "aws_ssoadmin_permission_set_invalid_description" {
  enabled = true
}

rule "aws_ssoadmin_permission_set_invalid_instance_arn" {
  enabled = true
}

rule "aws_ssoadmin_permission_set_invalid_name" {
  enabled = true
}

rule "aws_ssoadmin_permission_set_invalid_relay_state" {
  enabled = true
}

rule "aws_storagegateway_cache_invalid_disk_id" {
  enabled = true
}

rule "aws_storagegateway_cache_invalid_gateway_arn" {
  enabled = true
}

rule "aws_storagegateway_cached_iscsi_volume_invalid_gateway_arn" {
  enabled = true
}

rule "aws_storagegateway_cached_iscsi_volume_invalid_network_interface_id" {
  enabled = true
}

rule "aws_storagegateway_cached_iscsi_volume_invalid_snapshot_id" {
  enabled = true
}

rule "aws_storagegateway_cached_iscsi_volume_invalid_source_volume_arn" {
  enabled = true
}

rule "aws_storagegateway_cached_iscsi_volume_invalid_target_name" {
  enabled = true
}

rule "aws_storagegateway_file_system_association_invalid_audit_destination_arn" {
  enabled = true
}

rule "aws_storagegateway_file_system_association_invalid_gateway_arn" {
  enabled = true
}

rule "aws_storagegateway_file_system_association_invalid_location_arn" {
  enabled = true
}

rule "aws_storagegateway_file_system_association_invalid_password" {
  enabled = true
}

rule "aws_storagegateway_file_system_association_invalid_username" {
  enabled = true
}

rule "aws_storagegateway_gateway_invalid_activation_key" {
  enabled = true
}

rule "aws_storagegateway_gateway_invalid_gateway_name" {
  enabled = true
}

rule "aws_storagegateway_gateway_invalid_gateway_timezone" {
  enabled = true
}

rule "aws_storagegateway_gateway_invalid_gateway_type" {
  enabled = true
}

rule "aws_storagegateway_gateway_invalid_medium_changer_type" {
  enabled = true
}

rule "aws_storagegateway_gateway_invalid_smb_guest_password" {
  enabled = true
}

rule "aws_storagegateway_gateway_invalid_tape_drive_type" {
  enabled = true
}

rule "aws_storagegateway_nfs_file_share_invalid_default_storage_class" {
  enabled = true
}

rule "aws_storagegateway_nfs_file_share_invalid_gateway_arn" {
  enabled = true
}

rule "aws_storagegateway_nfs_file_share_invalid_kms_key_arn" {
  enabled = true
}

rule "aws_storagegateway_nfs_file_share_invalid_location_arn" {
  enabled = true
}

rule "aws_storagegateway_nfs_file_share_invalid_object_acl" {
  enabled = true
}

rule "aws_storagegateway_nfs_file_share_invalid_role_arn" {
  enabled = true
}

rule "aws_storagegateway_nfs_file_share_invalid_squash" {
  enabled = true
}

rule "aws_storagegateway_smb_file_share_invalid_authentication" {
  enabled = true
}

rule "aws_storagegateway_smb_file_share_invalid_default_storage_class" {
  enabled = true
}

rule "aws_storagegateway_smb_file_share_invalid_gateway_arn" {
  enabled = true
}

rule "aws_storagegateway_smb_file_share_invalid_kms_key_arn" {
  enabled = true
}

rule "aws_storagegateway_smb_file_share_invalid_location_arn" {
  enabled = true
}

rule "aws_storagegateway_smb_file_share_invalid_object_acl" {
  enabled = true
}

rule "aws_storagegateway_smb_file_share_invalid_role_arn" {
  enabled = true
}

rule "aws_storagegateway_stored_iscsi_volume_invalid_disk_id" {
  enabled = true
}

rule "aws_storagegateway_stored_iscsi_volume_invalid_gateway_arn" {
  enabled = true
}

rule "aws_storagegateway_stored_iscsi_volume_invalid_kms_key" {
  enabled = true
}

rule "aws_storagegateway_stored_iscsi_volume_invalid_network_interface_id" {
  enabled = true
}

rule "aws_storagegateway_stored_iscsi_volume_invalid_snapshot_id" {
  enabled = true
}

rule "aws_storagegateway_stored_iscsi_volume_invalid_target_name" {
  enabled = true
}

rule "aws_storagegateway_tape_pool_invalid_pool_name" {
  enabled = true
}

rule "aws_storagegateway_tape_pool_invalid_retention_lock_type" {
  enabled = true
}

rule "aws_storagegateway_tape_pool_invalid_storage_class" {
  enabled = true
}

rule "aws_storagegateway_upload_buffer_invalid_disk_id" {
  enabled = true
}

rule "aws_storagegateway_upload_buffer_invalid_gateway_arn" {
  enabled = true
}

rule "aws_storagegateway_working_storage_invalid_disk_id" {
  enabled = true
}

rule "aws_storagegateway_working_storage_invalid_gateway_arn" {
  enabled = true
}

rule "aws_swf_domain_invalid_description" {
  enabled = true
}

rule "aws_swf_domain_invalid_name" {
  enabled = true
}

rule "aws_swf_domain_invalid_workflow_execution_retention_period_in_days" {
  enabled = true
}

rule "aws_timestreamwrite_database_invalid_database_name" {
  enabled = true
}

rule "aws_timestreamwrite_database_invalid_kms_key_id" {
  enabled = true
}

rule "aws_timestreamwrite_table_invalid_database_name" {
  enabled = true
}

rule "aws_timestreamwrite_table_invalid_table_name" {
  enabled = true
}

rule "aws_transfer_access_invalid_external_id" {
  enabled = true
}

rule "aws_transfer_access_invalid_home_directory" {
  enabled = true
}

rule "aws_transfer_access_invalid_home_directory_type" {
  enabled = true
}

rule "aws_transfer_access_invalid_policy" {
  enabled = true
}

rule "aws_transfer_access_invalid_role" {
  enabled = true
}

rule "aws_transfer_access_invalid_server_id" {
  enabled = true
}

rule "aws_transfer_server_invalid_endpoint_type" {
  enabled = true
}

rule "aws_transfer_server_invalid_identity_provider_type" {
  enabled = true
}

rule "aws_transfer_server_invalid_invocation_role" {
  enabled = true
}

rule "aws_transfer_server_invalid_logging_role" {
  enabled = true
}

rule "aws_transfer_server_invalid_url" {
  enabled = true
}

rule "aws_transfer_ssh_key_invalid_body" {
  enabled = true
}

rule "aws_transfer_ssh_key_invalid_server_id" {
  enabled = true
}

rule "aws_transfer_ssh_key_invalid_user_name" {
  enabled = true
}

rule "aws_transfer_user_invalid_home_directory" {
  enabled = true
}

rule "aws_transfer_user_invalid_policy" {
  enabled = true
}

rule "aws_transfer_user_invalid_role" {
  enabled = true
}

rule "aws_transfer_user_invalid_server_id" {
  enabled = true
}

rule "aws_transfer_user_invalid_user_name" {
  enabled = true
}

rule "aws_vpc_endpoint_invalid_vpc_endpoint_type" {
  enabled = true
}

rule "aws_vpc_invalid_instance_tenancy" {
  enabled = true
}

rule "aws_vpc_ipam_pool_invalid_address_family" {
  enabled = true
}

rule "aws_vpc_ipam_pool_invalid_aws_service" {
  enabled = true
}

rule "aws_waf_byte_match_set_invalid_name" {
  enabled = true
}

rule "aws_waf_geo_match_set_invalid_name" {
  enabled = true
}

rule "aws_waf_ipset_invalid_name" {
  enabled = true
}

rule "aws_waf_rate_based_rule_invalid_metric_name" {
  enabled = true
}

rule "aws_waf_rate_based_rule_invalid_name" {
  enabled = true
}

rule "aws_waf_rate_based_rule_invalid_rate_key" {
  enabled = true
}

rule "aws_waf_regex_match_set_invalid_name" {
  enabled = true
}

rule "aws_waf_regex_pattern_set_invalid_name" {
  enabled = true
}

rule "aws_waf_rule_group_invalid_metric_name" {
  enabled = true
}

rule "aws_waf_rule_group_invalid_name" {
  enabled = true
}

rule "aws_waf_rule_invalid_metric_name" {
  enabled = true
}

rule "aws_waf_rule_invalid_name" {
  enabled = true
}

rule "aws_waf_size_constraint_set_invalid_name" {
  enabled = true
}

rule "aws_waf_sql_injection_match_set_invalid_name" {
  enabled = true
}

rule "aws_waf_web_acl_invalid_metric_name" {
  enabled = true
}

rule "aws_waf_web_acl_invalid_name" {
  enabled = true
}

rule "aws_waf_xss_match_set_invalid_name" {
  enabled = true
}

rule "aws_wafregional_byte_match_set_invalid_name" {
  enabled = true
}

rule "aws_wafregional_geo_match_set_invalid_name" {
  enabled = true
}

rule "aws_wafregional_ipset_invalid_name" {
  enabled = true
}

rule "aws_wafregional_rate_based_rule_invalid_metric_name" {
  enabled = true
}

rule "aws_wafregional_rate_based_rule_invalid_name" {
  enabled = true
}

rule "aws_wafregional_rate_based_rule_invalid_rate_key" {
  enabled = true
}

rule "aws_wafregional_regex_match_set_invalid_name" {
  enabled = true
}

rule "aws_wafregional_regex_pattern_set_invalid_name" {
  enabled = true
}

rule "aws_wafregional_rule_group_invalid_metric_name" {
  enabled = true
}

rule "aws_wafregional_rule_group_invalid_name" {
  enabled = true
}

rule "aws_wafregional_rule_invalid_metric_name" {
  enabled = true
}

rule "aws_wafregional_rule_invalid_name" {
  enabled = true
}

rule "aws_wafregional_size_constraint_set_invalid_name" {
  enabled = true
}

rule "aws_wafregional_sql_injection_match_set_invalid_name" {
  enabled = true
}

rule "aws_wafregional_web_acl_association_invalid_resource_arn" {
  enabled = true
}

rule "aws_wafregional_web_acl_association_invalid_web_acl_id" {
  enabled = true
}

rule "aws_wafregional_web_acl_invalid_metric_name" {
  enabled = true
}

rule "aws_wafregional_web_acl_invalid_name" {
  enabled = true
}

rule "aws_wafregional_xss_match_set_invalid_name" {
  enabled = true
}

rule "aws_wafv2_ip_set_invalid_description" {
  enabled = true
}

rule "aws_wafv2_ip_set_invalid_ip_address_version" {
  enabled = true
}

rule "aws_wafv2_ip_set_invalid_name" {
  enabled = true
}

rule "aws_wafv2_ip_set_invalid_scope" {
  enabled = true
}

rule "aws_wafv2_regex_pattern_set_invalid_description" {
  enabled = true
}

rule "aws_wafv2_regex_pattern_set_invalid_name" {
  enabled = true
}

rule "aws_wafv2_regex_pattern_set_invalid_scope" {
  enabled = true
}

rule "aws_wafv2_rule_group_invalid_description" {
  enabled = true
}

rule "aws_wafv2_rule_group_invalid_name" {
  enabled = true
}

rule "aws_wafv2_rule_group_invalid_scope" {
  enabled = true
}

rule "aws_wafv2_web_acl_association_invalid_resource_arn" {
  enabled = true
}

rule "aws_wafv2_web_acl_association_invalid_web_acl_arn" {
  enabled = true
}

rule "aws_wafv2_web_acl_invalid_description" {
  enabled = true
}

rule "aws_wafv2_web_acl_invalid_name" {
  enabled = true
}

rule "aws_wafv2_web_acl_invalid_scope" {
  enabled = true
}

rule "aws_wafv2_web_acl_logging_configuration_invalid_resource_arn" {
  enabled = true
}

rule "aws_worklink_fleet_invalid_audit_stream_arn" {
  enabled = true
}

rule "aws_worklink_fleet_invalid_device_ca_certificate" {
  enabled = true
}

rule "aws_worklink_fleet_invalid_display_name" {
  enabled = true
}

rule "aws_worklink_fleet_invalid_name" {
  enabled = true
}

rule "aws_worklink_website_certificate_authority_association_invalid_certificate" {
  enabled = true
}

rule "aws_worklink_website_certificate_authority_association_invalid_display_name" {
  enabled = true
}

rule "aws_worklink_website_certificate_authority_association_invalid_fleet_arn" {
  enabled = true
}

rule "aws_workspaces_directory_invalid_directory_id" {
  enabled = true
}

rule "aws_workspaces_workspace_invalid_bundle_id" {
  enabled = true
}

rule "aws_workspaces_workspace_invalid_directory_id" {
  enabled = true
}

rule "aws_workspaces_workspace_invalid_user_name" {
  enabled = true
}

rule "aws_xray_encryption_config_invalid_key_id" {
  enabled = true
}

rule "aws_xray_encryption_config_invalid_type" {
  enabled = true
}

rule "aws_xray_group_invalid_group_name" {
  enabled = true
}

rule "aws_xray_sampling_rule_invalid_host" {
  enabled = true
}

rule "aws_xray_sampling_rule_invalid_http_method" {
  enabled = true
}

rule "aws_xray_sampling_rule_invalid_resource_arn" {
  enabled = true
}

rule "aws_xray_sampling_rule_invalid_rule_name" {
  enabled = true
}

rule "aws_xray_sampling_rule_invalid_service_name" {
  enabled = true
}

rule "aws_xray_sampling_rule_invalid_service_type" {
  enabled = true
}

rule "aws_xray_sampling_rule_invalid_url_path" {
  enabled = true
}
