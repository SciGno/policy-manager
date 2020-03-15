// Tips
// schema.config().option('graph.schema_mode').set('Development')
// schema.config().option('graph.schema_mode').set('Production')
// schema.config().option('graph.schema_mode').get()
// schema.drop()
// schema.config().option('graph.allow_scan').set(true);
// schema.config().option('graph.allow_scan').set(false);
// schema.config().option('graph.allow_scan').get();

// Property Keys
schema.propertyKey('createdOn').Timestamp().single().ifNotExists().create()
schema.propertyKey('createdBy').Uuid().single().ifNotExists().create()
schema.propertyKey('modifiedOn').Timestamp().single().ifNotExists().create()
schema.propertyKey('modifiedBy').Uuid().single().ifNotExists().create()
schema.propertyKey('contact').Boolean().single().ifNotExists().create()
schema.propertyKey('verified').Boolean().single().ifNotExists().create()
schema.propertyKey('street').Text().single().ifNotExists().create()
schema.propertyKey('street2').Text().single().ifNotExists().create()
schema.propertyKey('city').Text().single().ifNotExists().create()
schema.propertyKey('state').Text().single().ifNotExists().create()
schema.propertyKey('country').Text().single().ifNotExists().create()
schema.propertyKey('postal').Text().single().ifNotExists().create()
schema.propertyKey('name').Text().single().ifNotExists().create()
schema.propertyKey('value').Text().single().ifNotExists().create()
schema.propertyKey('enabled').Boolean().single().ifNotExists().create()
schema.propertyKey('allow').Text().single().ifNotExists().create()
schema.propertyKey('deny').Text().single().ifNotExists().create()
schema.propertyKey('status').Text().single().ifNotExists().create()
schema.propertyKey('failures').Int().single().ifNotExists().create()
schema.propertyKey('avatar').Text().single().ifNotExists().create()
schema.propertyKey('statusUpdated').Timestamp().single().ifNotExists().create()
schema.propertyKey('sid').Text().single().ifNotExists().create()
schema.propertyKey('description').Text().single().ifNotExists().create()
schema.propertyKey('effect').Text().single().ifNotExists().create()
schema.propertyKey('resources').Text().single().ifNotExists().create()
schema.propertyKey('conditions').Text().single().ifNotExists().create()
schema.propertyKey('actions').Text().single().ifNotExists().create()
schema.propertyKey('values').Text().single().ifNotExists().create()

//===========
// Vertices
//===========
// User 
schema.propertyKey('userId').Uuid().single().ifNotExists().create()
schema.propertyKey('username').Text().single().ifNotExists().create()
schema.propertyKey('password').Text().single().ifNotExists().create()
schema.vertexLabel('user').partitionKey('userId').properties('username', 'password', 'status', 'failures', 'avatar', 'statusUpdated', 'createdOn', 'createdBy', 'modifiedOn', 'modifiedBy').ifNotExists().create()

// Access Token
schema.propertyKey('accessTokenId').Uuid().single().ifNotExists().create()
schema.propertyKey('accessTokenString').Text().single().ifNotExists().create()
schema.vertexLabel('accessToken').partitionKey('accessTokenId').properties('accessTokenString', 'createdOn', 'createdBy', 'modifiedOn', 'modifiedBy').ttl(3600).ifNotExists().create()

// Refresh Token
schema.propertyKey('refreshTokenId').Uuid().single().ifNotExists().create()
schema.propertyKey('refreshTokenString').Text().single().ifNotExists().create()
schema.vertexLabel('refreshToken').partitionKey('refreshTokenId').properties('refreshTokenString', 'createdOn', 'createdBy', 'modifiedOn', 'modifiedBy').ttl(2592000).ifNotExists().create()

// Email 
schema.propertyKey('emailId').Uuid().single().ifNotExists().create()
schema.propertyKey('emailAddress').Text().single().ifNotExists().create()
schema.vertexLabel('email').partitionKey('emailId').properties('emailAddress', 'contact', 'verified', 'createdOn', 'createdBy', 'modifiedOn', 'modifiedBy').ifNotExists().create()

// Phone 
schema.propertyKey('phoneId').Uuid().single().ifNotExists().create()
schema.propertyKey('number').Text().single().ifNotExists().create()
schema.vertexLabel('phone').partitionKey('phoneId').properties('number', 'contact', 'verified', 'name', 'createdOn', 'createdBy', 'modifiedOn', 'modifiedBy').ifNotExists().create()

// Address 
schema.propertyKey('addressId').Uuid().single().ifNotExists().create()
schema.vertexLabel('address').
partitionKey('addressId').properties('street', 'street2', 'city', 'state', 'country', 'postal', 'contact', 'verified', 'name', 'createdOn', 'createdBy', 'modifiedOn', 'modifiedBy').ifNotExists().create()

// Policy
schema.propertyKey('policyId').Uuid().single().ifNotExists().create()
schema.vertexLabel('policy').partitionKey('policyId').properties('name', 'description',  'createdOn', 'createdBy', 'modifiedOn', 'modifiedBy').ifNotExists().create()

// Statement
schema.propertyKey('statementId').Uuid().single().ifNotExists().create()
schema.vertexLabel('statement').partitionKey('statementId').properties('name', 'effect', 'createdOn', 'createdBy', 'modifiedOn', 'modifiedBy').ifNotExists().create()

// Actions
schema.propertyKey('actionId').Uuid().single().ifNotExists().create()
schema.vertexLabel('action').partitionKey('actionId').properties('name', 'actions', 'createdOn', 'createdBy', 'modifiedOn', 'modifiedBy').ifNotExists().create()

// Resource
schema.propertyKey('resourceId').Uuid().single().ifNotExists().create()
schema.vertexLabel('resource').partitionKey('resourceId').properties('name', 'values', 'createdOn', 'createdBy', 'modifiedOn', 'modifiedBy').ifNotExists().create()

// Conditions
schema.propertyKey('conditionId').Uuid().single().ifNotExists().create()
schema.vertexLabel('condition').partitionKey('conditionId').properties('name', 'values', 'createdOn', 'createdBy', 'modifiedOn', 'modifiedBy').ifNotExists().create()

// OU
schema.propertyKey('ouId').Uuid().single().ifNotExists().create()
schema.vertexLabel('ou').partitionKey('ouId').properties('name', 'createdOn', 'createdBy', 'modifiedOn', 'modifiedBy').ifNotExists().create()

//===========
// Edges
//===========
schema.edgeLabel('has').multiple().connection('user', 'email').ifNotExists().create()
schema.edgeLabel('has').multiple().connection('user', 'phone').ifNotExists().create()
schema.edgeLabel('has').multiple().connection('user', 'address').ifNotExists().create()
schema.edgeLabel('has').multiple().connection('user', 'accessToken').ttl(3600).ifNotExists().create()
schema.edgeLabel('has').multiple().connection('user', 'refreshToken').ttl(2592000).ifNotExists().create()
schema.edgeLabel('created').multiple().connection('user', 'policy').ifNotExists().create()
schema.edgeLabel('contains').multiple().connection('policy', 'statement').ifNotExists().create()
schema.edgeLabel('contains').multiple().connection('statement', 'action').ifNotExists().create()
schema.edgeLabel('contains').multiple().connection('statement', 'resource').ifNotExists().create()
schema.edgeLabel('contains').multiple().connection('statement', 'condition').ifNotExists().create()
schema.edgeLabel('applied').multiple().connection('policy', 'resource').ifNotExists().create()
schema.edgeLabel('applied').multiple().connection('policy', 'ou').ifNotExists().create()
schema.edgeLabel('belong').multiple().connection('user', 'ou').ifNotExists().create()

//===========
// Indexes
//===========
// User
schema.vertexLabel('user').index('byUsername').materialized().by('username').ifNotExists().add()
// Phone
schema.vertexLabel('phone').index('byName').materialized().by('name').ifNotExists().add()
// Email
schema.vertexLabel('email').index('byEmailAddress').materialized().by('emailAddress').ifNotExists().add()
// Address
schema.vertexLabel('address').index('byName').materialized().by('name').ifNotExists().add()
schema.vertexLabel('address').index('byCity').materialized().by('city').ifNotExists().add()
schema.vertexLabel('address').index('byState').materialized().by('state').ifNotExists().add()
schema.vertexLabel('address').index('byPostal').materialized().by('postal').ifNotExists().add()
// policy
schema.vertexLabel('policy').index('byName').materialized().by('name').ifNotExists().add()
// statement
schema.vertexLabel('statement').index('byName').materialized().by('name').ifNotExists().add()
// action
schema.vertexLabel('action').index('byName').materialized().by('name').ifNotExists().add()
// resource
schema.vertexLabel('resource').index('byName').materialized().by('name').ifNotExists().add()
// condition
schema.vertexLabel('condition').index('byName').materialized().by('name').ifNotExists().add()
// ou
schema.vertexLabel('ou').index('byName').materialized().by('name').ifNotExists().add()

'success'
