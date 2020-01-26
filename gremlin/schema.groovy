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
schema.propertyKey('tags').Text().single().ifNotExists().create()
schema.propertyKey('resource').Text().single().ifNotExists().create()
schema.propertyKey('summary').Text().single().ifNotExists().create()
schema.propertyKey('start').Timestamp().single().ifNotExists().create()
schema.propertyKey('end').Timestamp().single().ifNotExists().create()
schema.propertyKey('active').Boolean().single().ifNotExists().create()
schema.propertyKey('enabled').Boolean().single().ifNotExists().create()
schema.propertyKey('publishedOn').Timestamp().single().ifNotExists().create()
schema.propertyKey('rating').Int().single().ifNotExists().create()
schema.propertyKey('allow').Text().single().ifNotExists().create()
schema.propertyKey('deny').Text().single().ifNotExists().create()
schema.propertyKey('status').Text().single().ifNotExists().create()
schema.propertyKey('failures').Int().single().ifNotExists().create()
schema.propertyKey('avatar').Text().single().ifNotExists().create()
schema.propertyKey('logo').Text().single().ifNotExists().create()
schema.propertyKey('raw').Text().single().ifNotExists().create()
schema.propertyKey('statusUpdated').Timestamp().single().ifNotExists().create()
schema.propertyKey('description').Text().single().ifNotExists().create()
schema.propertyKey('subjects').Text().single().ifNotExists().create()
schema.propertyKey('effect').Text().single().ifNotExists().create()
schema.propertyKey('resources').Text().single().ifNotExists().create()
schema.propertyKey('conditions').Text().single().ifNotExists().create()
schema.propertyKey('actions').Text().single().ifNotExists().create()
schema.propertyKey('meta').Blob().single().ifNotExists().create()

//===========
// Vertices
//===========
// User 
schema.propertyKey('userId').Uuid().single().ifNotExists().create()
schema.propertyKey('username').Text().single().ifNotExists().create()
schema.propertyKey('password').Text().single().ifNotExists().create()
schema.vertexLabel('user').partitionKey('userId').properties('username', 'password', 'status', 'failures', 'avatar', 'statusUpdated', 'createdOn', 'createdBy', 'modifiedOn', 'modifiedBy').ifNotExists().create()

// Attribute
schema.propertyKey('attributeId').Uuid().single().ifNotExists().create()
schema.vertexLabel('attribute').partitionKey('attributeId').properties('name', 'value', 'createdOn', 'createdBy', 'modifiedOn', 'modifiedBy').ifNotExists().create()

// Publisher
schema.propertyKey('publisherId').Uuid().single().ifNotExists().create()
schema.vertexLabel('publisher').partitionKey('publisherId').properties('name', 'logo', 'status', 'verified', 'tags', 'createdOn', 'createdBy', 'modifiedOn', 'modifiedBy').ifNotExists().create()

// Access Token
schema.propertyKey('accessTokenId').Uuid().single().ifNotExists().create()
schema.propertyKey('accessTokenString').Text().single().ifNotExists().create()
schema.vertexLabel('accessToken').partitionKey('accessTokenId').properties('accessTokenString', 'createdOn', 'createdBy', 'modifiedOn', 'modifiedBy').ttl(3600).ifNotExists().create()

// Refresh Token
schema.propertyKey('refreshTokenId').Uuid().single().ifNotExists().create()
schema.propertyKey('refreshTokenString').Text().single().ifNotExists().create()
schema.vertexLabel('refreshToken').partitionKey('refreshTokenId').properties('refreshTokenString', 'createdOn', 'createdBy', 'modifiedOn', 'modifiedBy').ttl(2592000).ifNotExists().create()

// Policy
schema.propertyKey('policyId').Uuid().single().ifNotExists().create()
schema.vertexLabel('policy').partitionKey('policyId').properties('name', 'description', 'subjects', 'effect', 'actions', 'enabled', 'resources', 'conditions', 'meta', 'createdOn', 'createdBy', 'modifiedOn', 'modifiedBy').ifNotExists().create()

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

// Campaign 
schema.propertyKey('campaignId').Uuid().single().ifNotExists().create()
schema.vertexLabel('campaign').partitionKey('campaignId').properties('name', 'summary', 'start', 'end', 'active', 'createdOn', 'createdBy', 'modifiedOn', 'modifiedBy').ifNotExists().create()

// Publication 
schema.propertyKey('publicationId').Uuid().single().ifNotExists().create()
schema.vertexLabel('publication').partitionKey('publicationId').properties('name', 'summary', 'active', 'resource', 'tags', 'start', 'end', 'createdOn', 'createdBy', 'modifiedOn', 'modifiedBy').ifNotExists().create()

// Tag 
schema.propertyKey('tagId').Uuid().single().ifNotExists().create()
schema.vertexLabel('tag').partitionKey('campaignId').properties('name', 'createdOn', 'createdBy', 'modifiedOn', 'modifiedBy').ifNotExists().create()

//===========
// Edges
//===========
schema.edgeLabel('has').multiple().connection('user', 'attribute').ifNotExists().create()
schema.edgeLabel('has').multiple().connection('user', 'email').ifNotExists().create()
schema.edgeLabel('has').multiple().connection('user', 'phone').ifNotExists().create()
schema.edgeLabel('has').multiple().connection('user', 'address').ifNotExists().create()
schema.edgeLabel('has').multiple().connection('user', 'accessToken').ttl(3600).ifNotExists().create()
schema.edgeLabel('has').multiple().connection('user', 'refreshToken').ttl(2592000).ifNotExists().create()
schema.edgeLabel('registered').multiple().connection('user', 'publisher').ifNotExists().create()
schema.edgeLabel('subscribed').multiple().connection('user', 'publisher').ifNotExists().create()
schema.edgeLabel('viewed').multiple().connection('user', 'publication').ifNotExists().create()
// --------
schema.edgeLabel('rated').multiple().properties('rating').connection('user', 'publication').ifNotExists().create()
// Publisher -> Email 
schema.edgeLabel('has').multiple().connection('publisher', 'email').ifNotExists().create()
// Publisher -> Phone 
schema.edgeLabel('has').multiple().connection('publisher', 'phone').ifNotExists().create()
// Publisher -> Address 
schema.edgeLabel('has').multiple().connection('publisher', 'address').ifNotExists().create()
// Publisher -> Policy 
schema.edgeLabel('created').multiple().connection('publisher', 'policy').ifNotExists().create()
// Policy -> User
schema.edgeLabel('assigned').multiple().connection('policy', 'user').ifNotExists().create()
// Campaign -> Attribute 
schema.edgeLabel('targets').multiple().connection('campaign', 'attribute').ifNotExists().create()
// Publication -> Attribute 
schema.edgeLabel('uses').multiple().connection('publication', 'attribute').ifNotExists().create()
// Publication -> User 
schema.edgeLabel('for').multiple().connection('publication', 'user').ifNotExists().create()
// Publisher -> Campaign 
schema.edgeLabel('created').multiple().connection('publisher', 'campaign').ifNotExists().create()
schema.edgeLabel('published').multiple().connection('publisher', 'campaign').ifNotExists().create()
// Campaign -> Publication 
schema.edgeLabel('has').multiple().connection('campaign', 'publication').ifNotExists().create()
// Publisher -> Publication
schema.edgeLabel('published').multiple().connection('publisher', 'publication').ifNotExists().create()
schema.edgeLabel('created').multiple().connection('publisher', 'publication').ifNotExists().create()
// Publisher -> Tag
schema.edgeLabel('has').multiple().connection('publisher', 'tag').ifNotExists().create()
// Publication -> Tag
schema.edgeLabel('has').multiple().connection('publication', 'tag').ifNotExists().create()

//===========
// Indexes.
//===========
// User
schema.vertexLabel('user').index('byUsername').materialized().by('username').ifNotExists().add()
// Attribute
schema.vertexLabel('attribute').index('byName').materialized().by('name').ifNotExists().add()
// Phone
schema.vertexLabel('phone').index('byName').materialized().by('name').ifNotExists().add()
// Email
schema.vertexLabel('email').index('byEmailAddress').materialized().by('emailAddress').ifNotExists().add()
// Address
schema.vertexLabel('address').index('byName').materialized().by('name').ifNotExists().add()
schema.vertexLabel('address').index('byCity').materialized().by('city').ifNotExists().add()
schema.vertexLabel('address').index('byState').materialized().by('state').ifNotExists().add()
schema.vertexLabel('address').index('byPostal').materialized().by('postal').ifNotExists().add()
// publisher
schema.vertexLabel('publisher').index('byName').materialized().by('name').ifNotExists().add()
// publication
schema.vertexLabel('publication').index('byName').materialized().by('name').ifNotExists().add()
// campaign
schema.vertexLabel('campaign').index('byName').materialized().by('name').ifNotExists().add()
// policy
schema.vertexLabel('policy').index('byName').materialized().by('name').ifNotExists().add()
// tag
schema.vertexLabel('tag').index('byName').materialized().by('name').ifNotExists().add()


'success'
