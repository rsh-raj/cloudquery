# Table: stripe_treasury_credit_reversals

https://stripe.com/docs/api/treasury_credit_reversals

The primary key for this table is **id**.

## Relations

This table depends on [stripe_treasury_financial_accounts](stripe_treasury_financial_accounts.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|amount|Int|
|created|Timestamp|
|currency|String|
|financial_account|String|
|hosted_regulatory_receipt_url|String|
|livemode|Bool|
|metadata|JSON|
|network|String|
|object|String|
|received_credit|String|
|status|String|
|status_transitions|JSON|
|transaction|JSON|