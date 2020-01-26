

g.V().hasLabel('user').project('username', 'userId', 'emailAddress').by(values('username')).by(values('userId')).by(out("has").hasLabel('email').values("emailAddress").fold())
{
  "username": "llopez",
  "userId": "dd585dad-9f7e-4f69-8b17-05803d5c14b0",
  "emailAddress": [
    "leandro@scigno.com",
    "leandro.lopez@gmail.com"
  ]
}

// Get emails for a particular user
g.V().has('user', 'username', 'llopez').as('user').outE('createdEmail').inV().as('email').select('user', 'email').by('username').by('emailAddress')
// OUTPUT
[
 {
  "email": "leandro.lopez@gmail.com",
  "user": "llopez"
 },
 {
  "email": "leandro@scigno.com",
  "user": "llopez"
 }
]

// Get emails for a particular user
g.V().has('user', 'username', 'llopez').as('user').outE('createdEmail').inV().as('email').select('user', 'email').fold()
[
  {
    "user": {
      "id": "{~label=user, userId=71e5ed34-eb88-4ff4-b39c-851ebf61b337}",
      "label": "user",
      "type": "vertex",
      "properties": {
        "modifiedOn": [
          {
            "id": "{~label=modifiedOn, ~out_vertex={~label=user, userId=71e5ed34-eb88-4ff4-b39c-851ebf61b337}, ~local_id=00000000-0000-803f-0000-000000000000}",
            "value": "2019-04-14T00:20:37.375Z"
          }
        ],
        "createdBy": [
          {
            "id": "{~label=createdBy, ~out_vertex={~label=user, userId=71e5ed34-eb88-4ff4-b39c-851ebf61b337}, ~local_id=00000000-0000-803e-0000-000000000000}",
            "value": "71e5ed34-eb88-4ff4-b39c-851ebf61b337"
          }
        ],
        "modifiedBy": [
          {
            "id": "{~label=modifiedBy, ~out_vertex={~label=user, userId=71e5ed34-eb88-4ff4-b39c-851ebf61b337}, ~local_id=00000000-0000-8040-0000-000000000000}",
            "value": "71e5ed34-eb88-4ff4-b39c-851ebf61b337"
          }
        ],
        "userId": [
          {
            "id": "{~label=userId, ~out_vertex={~label=user, userId=71e5ed34-eb88-4ff4-b39c-851ebf61b337}, ~local_id=00000000-0000-0000-0000-000000000000}",
            "value": "71e5ed34-eb88-4ff4-b39c-851ebf61b337"
          }
        ],
        "createdOn": [
          {
            "id": "{~label=createdOn, ~out_vertex={~label=user, userId=71e5ed34-eb88-4ff4-b39c-851ebf61b337}, ~local_id=00000000-0000-803d-0000-000000000000}",
            "value": "2019-04-14T00:20:37.375Z"
          }
        ],
        "username": [
          {
            "id": "{~label=username, ~out_vertex={~label=user, userId=71e5ed34-eb88-4ff4-b39c-851ebf61b337}, ~local_id=00000000-0000-8058-0000-000000000000}",
            "value": "llopez"
          }
        ]
      }
    },
    "email": {
      "id": "{~label=email, emailId=a934c36e-b6ac-4da6-904c-008bfdf707ba}",
      "label": "email",
      "type": "vertex",
      "properties": {
        "modifiedOn": [
          {
            "id": "{~label=modifiedOn, ~out_vertex={~label=email, emailId=a934c36e-b6ac-4da6-904c-008bfdf707ba}, ~local_id=00000000-0000-803f-0000-000000000000}",
            "value": "2019-04-14T00:20:37.376Z"
          }
        ],
        "emailAddress": [
          {
            "id": "{~label=emailAddress, ~out_vertex={~label=email, emailId=a934c36e-b6ac-4da6-904c-008bfdf707ba}, ~local_id=00000000-0000-8063-0000-000000000000}",
            "value": "leandro.lopez@gmail.com"
          }
        ],
        "createdBy": [
          {
            "id": "{~label=createdBy, ~out_vertex={~label=email, emailId=a934c36e-b6ac-4da6-904c-008bfdf707ba}, ~local_id=00000000-0000-803e-0000-000000000000}",
            "value": "71e5ed34-eb88-4ff4-b39c-851ebf61b337"
          }
        ],
        "emailId": [
          {
            "id": "{~label=emailId, ~out_vertex={~label=email, emailId=a934c36e-b6ac-4da6-904c-008bfdf707ba}, ~local_id=00000000-0000-0000-0000-000000000000}",
            "value": "a934c36e-b6ac-4da6-904c-008bfdf707ba"
          }
        ],
        "modifiedBy": [
          {
            "id": "{~label=modifiedBy, ~out_vertex={~label=email, emailId=a934c36e-b6ac-4da6-904c-008bfdf707ba}, ~local_id=00000000-0000-8040-0000-000000000000}",
            "value": "71e5ed34-eb88-4ff4-b39c-851ebf61b337"
          }
        ],
        "createdOn": [
          {
            "id": "{~label=createdOn, ~out_vertex={~label=email, emailId=a934c36e-b6ac-4da6-904c-008bfdf707ba}, ~local_id=00000000-0000-803d-0000-000000000000}",
            "value": "2019-04-14T00:20:37.376Z"
          }
        ]
      }
    }
  },
  {
    "user": {
      "id": "{~label=user, userId=71e5ed34-eb88-4ff4-b39c-851ebf61b337}",
      "label": "user",
      "type": "vertex",
      "properties": {
        "modifiedOn": [
          {
            "id": "{~label=modifiedOn, ~out_vertex={~label=user, userId=71e5ed34-eb88-4ff4-b39c-851ebf61b337}, ~local_id=00000000-0000-803f-0000-000000000000}",
            "value": "2019-04-14T00:20:37.375Z"
          }
        ],
        "createdBy": [
          {
            "id": "{~label=createdBy, ~out_vertex={~label=user, userId=71e5ed34-eb88-4ff4-b39c-851ebf61b337}, ~local_id=00000000-0000-803e-0000-000000000000}",
            "value": "71e5ed34-eb88-4ff4-b39c-851ebf61b337"
          }
        ],
        "modifiedBy": [
          {
            "id": "{~label=modifiedBy, ~out_vertex={~label=user, userId=71e5ed34-eb88-4ff4-b39c-851ebf61b337}, ~local_id=00000000-0000-8040-0000-000000000000}",
            "value": "71e5ed34-eb88-4ff4-b39c-851ebf61b337"
          }
        ],
        "userId": [
          {
            "id": "{~label=userId, ~out_vertex={~label=user, userId=71e5ed34-eb88-4ff4-b39c-851ebf61b337}, ~local_id=00000000-0000-0000-0000-000000000000}",
            "value": "71e5ed34-eb88-4ff4-b39c-851ebf61b337"
          }
        ],
        "createdOn": [
          {
            "id": "{~label=createdOn, ~out_vertex={~label=user, userId=71e5ed34-eb88-4ff4-b39c-851ebf61b337}, ~local_id=00000000-0000-803d-0000-000000000000}",
            "value": "2019-04-14T00:20:37.375Z"
          }
        ],
        "username": [
          {
            "id": "{~label=username, ~out_vertex={~label=user, userId=71e5ed34-eb88-4ff4-b39c-851ebf61b337}, ~local_id=00000000-0000-8058-0000-000000000000}",
            "value": "llopez"
          }
        ]
      }
    },
    "email": {
      "id": "{~label=email, emailId=fb2fc680-ce71-4a98-992d-29a4020d2caf}",
      "label": "email",
      "type": "vertex",
      "properties": {
        "modifiedOn": [
          {
            "id": "{~label=modifiedOn, ~out_vertex={~label=email, emailId=fb2fc680-ce71-4a98-992d-29a4020d2caf}, ~local_id=00000000-0000-803f-0000-000000000000}",
            "value": "2019-04-14T00:20:37.377Z"
          }
        ],
        "emailAddress": [
          {
            "id": "{~label=emailAddress, ~out_vertex={~label=email, emailId=fb2fc680-ce71-4a98-992d-29a4020d2caf}, ~local_id=00000000-0000-8063-0000-000000000000}",
            "value": "leandro@scigno.com"
          }
        ],
        "createdBy": [
          {
            "id": "{~label=createdBy, ~out_vertex={~label=email, emailId=fb2fc680-ce71-4a98-992d-29a4020d2caf}, ~local_id=00000000-0000-803e-0000-000000000000}",
            "value": "71e5ed34-eb88-4ff4-b39c-851ebf61b337"
          }
        ],
        "emailId": [
          {
            "id": "{~label=emailId, ~out_vertex={~label=email, emailId=fb2fc680-ce71-4a98-992d-29a4020d2caf}, ~local_id=00000000-0000-0000-0000-000000000000}",
            "value": "fb2fc680-ce71-4a98-992d-29a4020d2caf"
          }
        ],
        "modifiedBy": [
          {
            "id": "{~label=modifiedBy, ~out_vertex={~label=email, emailId=fb2fc680-ce71-4a98-992d-29a4020d2caf}, ~local_id=00000000-0000-8040-0000-000000000000}",
            "value": "71e5ed34-eb88-4ff4-b39c-851ebf61b337"
          }
        ],
        "createdOn": [
          {
            "id": "{~label=createdOn, ~out_vertex={~label=email, emailId=fb2fc680-ce71-4a98-992d-29a4020d2caf}, ~local_id=00000000-0000-803d-0000-000000000000}",
            "value": "2019-04-14T00:20:37.377Z"
          }
        ]
      }
    }
  }
]