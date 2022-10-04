---
page_title: "genesyscloud_outbound_ruleset Resource - terraform-provider-genesyscloud"
subcategory: ""
description: |-
  Genesys Cloud outbound ruleset
---
# genesyscloud_outbound_ruleset (Resource)

Genesys Cloud outbound ruleset

## API Usage
The following Genesys Cloud APIs are used by this resource. Ensure your OAuth Client has been granted the necessary scopes and permissions to perform these operations:

* [POST /api/v2/outbound/rulesets](https://developer.genesys.cloud/devapps/api-explorer#post-api-v2-outbound-rulesets)
* [GET /api/v2/outbound/rulesets/{ruleSetId}](https://developer.genesys.cloud/devapps/api-explorer#get-api-v2-outbound-rulesets--ruleSetId-)
* [GET /api/v2/outbound/rulesets](https://developer.genesys.cloud/devapps/api-explorer#get-api-v2-outbound-rulesets)
* [DELETE /api/v2/outbound/rulesets/{ruleSetId}](https://developer.genesys.cloud/devapps/api-explorer#delete-api-v2-outbound-rulesets--ruleSetId-)
* [PUT /api/v2/outbound/rulesets/{ruleSetId}](https://developer.genesys.cloud/devapps/api-explorer#put-api-v2-outbound-rulesets--ruleSetId-)

## Example Usage

```terraform
resource "genesyscloud_outbound_ruleset" "example_outbound_ruleset" {
  name            = ""
  contact_list_id = genesyscloud_outbound_contact_list.contact_list.id
  queue_id        = genesyscloud_routing_queue.queue.id
  rules {
    name     = ""
    order    = 0
    category = "DIALER_PRECALL" // Possible values: DIALER_PRECALL, DIALER_WRAPUP
    conditions {
      type                       = "wrapupCondition" // Possible values: wrapupCondition, systemDispositionCondition, contactAttributeCondition, phoneNumberCondition, phoneNumberTypeCondition, callAnalysisCondition, contactPropertyCondition, dataActionCondition
      inverted                   = true
      attribute_name             = ""
      value                      = ""
      value_type                 = "STRING" // Possible values: STRING, NUMERIC, DATETIME, PERIOD
      operator                   = "EQUALS" // Possible values: EQUALS, LESS_THAN, LESS_THAN_EQUALS, GREATER_THAN, GREATER_THAN_EQUALS, CONTAINS, BEGINS_WITH, ENDS_WITH, BEFORE, AFTER, IN
      codes                      = []
      property                   = ""
      property_type              = "LAST_ATTEMPT_BY_COLUMN" // Possible values: LAST_ATTEMPT_BY_COLUMN, LAST_ATTEMPT_OVERALL, LAST_WRAPUP_BY_COLUMN, LAST_WRAPUP_OVERALL
      data_action_id             = genesyscloud_integration_action.data_action.id
      data_not_found_resolution  = true
      contact_id_field           = ""
      call_analysis_result_field = ""
      agent_wrapup_field         = ""
      contact_column_to_data_action_field_mappings {
        contact_column_name = ""
        data_action_field   = ""
      }
      predicates {
        output_field                    = ""
        output_operator                 = "EQUALS" // Possible values: EQUALS, LESS_THAN, LESS_THAN_EQUALS, GREATER_THAN, GREATER_THAN_EQUALS, CONTAINS, BEGINS_WITH, ENDS_WITH, BEFORE, AFTER
        comparison_value                = ""
        inverted                        = true
        output_field_missing_resolution = true
      }
    }
    actions {
      type             = "Action"      // Possible values: Action, modifyContactAttribute, dataActionBehavior
      action_type_name = "DO_NOT_DIAL" // Possible values: DO_NOT_DIAL, MODIFY_CONTACT_ATTRIBUTE, SWITCH_TO_PREVIEW, APPEND_NUMBER_TO_DNC_LIST, SCHEDULE_CALLBACK, CONTACT_UNCALLABLE, NUMBER_UNCALLABLE, SET_CALLER_ID, SET_SKILLS, DATA_ACTION
      update_option    = "SET"         // Possible values: SET, INCREMENT, DECREMENT, CURRENT_TIME
      properties       = {}
      data_action_id   = genesyscloud_integration_action.data_action.id
      contact_column_to_data_action_field_mappings {
        contact_column_name = ""
        data_action_field   = ""
      }
      contact_id_field           = ""
      call_analysis_result_field = ""
      agent_wrapup_field         = ""
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the RuleSet.

### Optional

- `contact_list_id` (String) A ContactList to provide user-interface suggestions for contact columns on relevant conditions and actions.
- `queue_id` (String) A Queue to provide user-interface suggestions for wrap-up codes on relevant conditions and actions.
- `rules` (Block List) The list of rules. (see [below for nested schema](#nestedblock--rules))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--rules"></a>
### Nested Schema for `rules`

Required:

- `category` (String) The category of the rule.
- `conditions` (Block List, Min: 1) A list of Conditions. All of the Conditions must evaluate to true to trigger the actions. (see [below for nested schema](#nestedblock--rules--conditions))
- `name` (String) The name of the rule.

Optional:

- `actions` (Block List) The list of actions to be taken if the conditions are true. (see [below for nested schema](#nestedblock--rules--actions))
- `order` (Number) The ranked order of the rule. Rules are processed from lowest number to highest.

<a id="nestedblock--rules--conditions"></a>
### Nested Schema for `rules.conditions`

Optional:

- `agent_wrapup_field` (String) The input field from the data action that the agentWrapup will be passed to for this condition. Valid for a wrapup dataActionCondition.
- `attribute_name` (String) An attribute name associated with this Condition. Required for a contactAttributeCondition.
- `call_analysis_result_field` (String) The input field from the data action that the callAnalysisResult will be passed to for this condition. Valid for a wrapup dataActionCondition.
- `codes` (List of String) List of wrap-up code identifiers. Required for a wrapupCondition.
- `contact_column_to_data_action_field_mappings` (Block Set) A list of mappings defining which contact data fields will be passed to which data action input fields for this condition. Valid for a dataActionCondition. (see [below for nested schema](#nestedblock--rules--conditions--contact_column_to_data_action_field_mappings))
- `contact_id_field` (String) The input field from the data action that the contactId will be passed to for this condition. Valid for a dataActionCondition.
- `data_action_id` (String) The Data Action to use for this condition. Required for a dataActionCondition.
- `data_not_found_resolution` (Boolean) The result of this condition if the data action returns a result indicating there was no data. Required for a DataActionCondition.
- `inverted` (Boolean) If true, inverts the result of evaluating this Condition. Default is false.
- `operator` (String) An operation with which to evaluate the Condition. Not used for a DataActionCondition.
- `predicates` (Block Set) A list of predicates defining the comparisons to use for this condition. Required for a dataActionCondition. (see [below for nested schema](#nestedblock--rules--conditions--predicates))
- `property` (String) A value associated with the property type of this Condition. Required for a contactPropertyCondition.
- `property_type` (String) The type of the property associated with this Condition. Required for a contactPropertyCondition.
- `type` (String) The type of the condition.
- `value` (String) A value associated with this Condition. This could be text, a number, or a relative time. Not used for a DataActionCondition.
- `value_type` (String) The type of the value associated with this Condition. Not used for a DataActionCondition.

<a id="nestedblock--rules--conditions--contact_column_to_data_action_field_mappings"></a>
### Nested Schema for `rules.conditions.contact_column_to_data_action_field_mappings`

Required:

- `contact_column_name` (String) The name of a contact column whose data will be passed to the data action
- `data_action_field` (String) The name of an input field from the data action that the contact column data will be passed to


<a id="nestedblock--rules--conditions--predicates"></a>
### Nested Schema for `rules.conditions.predicates`

Required:

- `comparison_value` (String) The value to compare against for this condition
- `inverted` (Boolean) If true, inverts the result of evaluating this Predicate. Default is false.
- `output_field` (String) The name of an output field from the data action's output to use for this condition
- `output_field_missing_resolution` (Boolean) The result of this predicate if the requested output field is missing from the data action's result
- `output_operator` (String) The operation with which to evaluate this condition



<a id="nestedblock--rules--actions"></a>
### Nested Schema for `rules.actions`

Required:

- `action_type_name` (String) Additional type specification for this DialerAction.
- `type` (String) The type of this DialerAction.

Optional:

- `agent_wrapup_field` (String) The input field from the data action that the agentWrapup will be passed to for this condition. Valid for a wrapup dataActionBehavior.
- `call_analysis_result_field` (String) The input field from the data action that the callAnalysisResult will be passed to for this condition. Valid for a wrapup dataActionBehavior.
- `contact_column_to_data_action_field_mappings` (Block Set) A list of mappings defining which contact data fields will be passed to which data action input fields for this condition. Valid for a dataActionBehavior. (see [below for nested schema](#nestedblock--rules--actions--contact_column_to_data_action_field_mappings))
- `contact_id_field` (String) The input field from the data action that the contactId will be passed to for this condition. Valid for a dataActionBehavior.
- `data_action_id` (String) The Data Action to use for this action. Required for a dataActionBehavior.
- `properties` (Map of String) A map of key-value pairs pertinent to the DialerAction. Different types of DialerActions require different properties. MODIFY_CONTACT_ATTRIBUTE with an updateOption of SET takes a contact column as the key and accepts any value. SCHEDULE_CALLBACK takes a key 'callbackOffset' that specifies how far in the future the callback should be scheduled, in minutes. SET_CALLER_ID takes two keys: 'callerAddress', which should be the caller id phone number, and 'callerName'. For either key, you can also specify a column on the contact to get the value from. To do this, specify 'contact.Column', where 'Column' is the name of the contact column from which to get the value. SET_SKILLS takes a key 'skills' with an array of skill ids wrapped into a string (Example: {'skills': '['skillIdHere']'} ).
- `update_option` (String) Specifies how a contact attribute should be updated. Required for MODIFY_CONTACT_ATTRIBUTE.

<a id="nestedblock--rules--actions--contact_column_to_data_action_field_mappings"></a>
### Nested Schema for `rules.actions.contact_column_to_data_action_field_mappings`

Required:

- `contact_column_name` (String) The name of a contact column whose data will be passed to the data action
- `data_action_field` (String) The name of an input field from the data action that the contact column data will be passed to
