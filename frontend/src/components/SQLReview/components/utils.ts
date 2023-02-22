import {
  convertPolicyRuleToRuleTemplate,
  RuleConfigComponent,
  RuleLevel,
  RuleTemplate,
  RuleType,
  SQLReviewPolicy,
  TEMPLATE_LIST,
} from "@/types";
import { PayloadValueType } from "./RuleConfigComponents";

export const rulesToTemplate = (
  review: SQLReviewPolicy,
  withDisabled = false
) => {
  const ruleTemplateList: RuleTemplate[] = [];
  const ruleTemplateMap: Map<RuleType, RuleTemplate> = TEMPLATE_LIST.reduce(
    (map, template) => {
      for (const rule of template.ruleList) {
        map.set(rule.type, rule);
      }
      return map;
    },
    new Map<RuleType, RuleTemplate>()
  );

  for (const policyRule of review.ruleList) {
    if (policyRule.level === RuleLevel.DISABLED && !withDisabled) {
      continue;
    }

    const rule = ruleTemplateMap.get(policyRule.type);
    if (!rule) {
      continue;
    }

    const data = convertPolicyRuleToRuleTemplate(policyRule, rule);
    if (data) {
      ruleTemplateList.push(data);
    }
    ruleTemplateMap.delete(policyRule.type);
  }

  if (withDisabled) {
    for (const rule of ruleTemplateMap.values()) {
      ruleTemplateList.push({
        ...rule,
        level: RuleLevel.DISABLED,
      });
    }
  }

  return {
    id: `bb.sql-review.environment-policy.${review.environment.resourceId}`,
    review,
    ruleList: ruleTemplateList,
  };
};

export const payloadValueListToComponentList = (
  rule: RuleTemplate,
  data: PayloadValueType[]
) => {
  const componentList = rule.componentList.reduce<RuleConfigComponent[]>(
    (list, component, index) => {
      switch (component.payload.type) {
        case "STRING_ARRAY":
          list.push({
            ...component,
            payload: {
              ...component.payload,
              value: data[index] as string[],
            },
          });
          break;
        case "NUMBER":
          list.push({
            ...component,
            payload: {
              ...component.payload,
              value: data[index] as number,
            },
          });
          break;
        case "BOOLEAN":
          list.push({
            ...component,
            payload: {
              ...component.payload,
              value: data[index] as boolean,
            },
          });
          break;
        default:
          list.push({
            ...component,
            payload: {
              ...component.payload,
              value: data[index] as string,
            },
          });
          break;
      }
      return list;
    },
    []
  );
  return componentList;
};
