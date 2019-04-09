package schema

import (
	"testing"

	"github.com/snapiz/go-vue-starter/packages/tgo"
	"github.com/snapiz/go-vue-starter/services/main/src/models"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/graphql-go/graphql"
)

func TestUpdateUser_AnonymousAccessShouldBeDenied(t *testing.T) {
	query := `
        mutation UpdateUserMutation($input: UpdateUserInput!) {
          updateUser(input: $input) {
			user {
				displayName
				picture
			}
			clientMutationId
		  }
        }
	  `
	params := map[string]interface{}{
		"input": map[string]interface{}{
			"displayName":      "Anonymous",
			"picture":          "Anonymous",
			"clientMutationId": "abcde",
		},
	}

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"updateUser": nil,
		},
	}

	result := graphql.Do(graphql.Params{
		Schema:         Schema,
		RequestString:  query,
		Context:        tgo.Anonymous,
		VariableValues: params,
	})

	assert.Equal(t, expected.Data, result.Data, "User should be nil")

	if assert.NotNil(t, result.Errors) {
		assert.Equal(t, "Anonymous access is denied", result.Errors[0].Message, "Anonymous should not update")
	}
}

func TestUpdateUser_DisplayNameShouldBeInvalid(t *testing.T) {
	query := `
        mutation UpdateUserMutation($input: UpdateUserInput!) {
          updateUser(input: $input) {
			user {
				displayName
				picture
			}
			clientMutationId
		  }
        }
	  `
	params := map[string]interface{}{
		"input": map[string]interface{}{
			"displayName":      "Jo",
			"picture":          "",
			"clientMutationId": "abcde",
		},
	}

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"updateUser": nil,
		},
	}

	result := graphql.Do(graphql.Params{
		Schema:         Schema,
		RequestString:  query,
		Context:        tgo.John,
		VariableValues: params,
	})

	assert.Equal(t, expected.Data, result.Data, "User should be nil")

	if assert.NotNil(t, result.Errors) {
		assert.Equal(t, "The displayName field must be between 3 and 50 characters long", result.Errors[0].Message, "should not update invalid displayName")
	}

	query = `
        mutation UpdateUserMutation($input: UpdateUserInput!) {
          updateUser(input: $input) {
			user {
				displayName
				picture
			}
			clientMutationId
		  }
        }
	  `
	params = map[string]interface{}{
		"input": map[string]interface{}{
			"displayName":      "John Doe John Doe John Doe John Doe John Doe John D",
			"picture":          "",
			"clientMutationId": "abcde",
		},
	}

	expected = &graphql.Result{
		Data: map[string]interface{}{
			"updateUser": nil,
		},
	}

	result = graphql.Do(graphql.Params{
		Schema:         Schema,
		RequestString:  query,
		Context:        tgo.John,
		VariableValues: params,
	})

	assert.Equal(t, expected.Data, result.Data, "User should be nil")

	if assert.NotNil(t, result.Errors) {
		assert.Equal(t, "The displayName field must be between 3 and 50 characters long", result.Errors[0].Message, "should not update invalid displayName")
	}
}

func TestUpdateUser_PictureShouldBeInvalid(t *testing.T) {
	query := `
        mutation UpdateUserMutation($input: UpdateUserInput!) {
          updateUser(input: $input) {
			user {
				displayName
				picture
			}
			clientMutationId
		  }
        }
	  `
	params := map[string]interface{}{
		"input": map[string]interface{}{
			"displayName":      "John doe",
			"picture":          "not_valid_url",
			"clientMutationId": "abcde",
		},
	}

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"updateUser": nil,
		},
	}

	result := graphql.Do(graphql.Params{
		Schema:         Schema,
		RequestString:  query,
		Context:        tgo.John,
		VariableValues: params,
	})

	assert.Equal(t, expected.Data, result.Data, "User should be nil")

	if assert.NotNil(t, result.Errors) {
		assert.Equal(t, "The picture field is invalid URL", result.Errors[0].Message, "should not update invalid picture")
	}
}

func TestUpdateUser_UsernameAlreadyExists(t *testing.T) {
	query := `
        mutation UpdateUserMutation($input: UpdateUserInput!) {
          updateUser(input: $input) {
			user {
				displayName
				picture
			}
			clientMutationId
		  }
        }
	  `
	params := map[string]interface{}{
		"input": map[string]interface{}{
			"username":         "albert",
			"displayName":      "John doe",
			"picture":          "",
			"clientMutationId": "abcde",
		},
	}

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"updateUser": nil,
		},
	}

	result := graphql.Do(graphql.Params{
		Schema:         Schema,
		RequestString:  query,
		Context:        tgo.John,
		VariableValues: params,
	})

	assert.Equal(t, expected.Data, result.Data, "User should be nil")

	if assert.NotNil(t, result.Errors) {
		assert.Equal(t, "errors.auth.usernameAlreadyExists", result.Errors[0].Message, "should not update invalid username")
	}
}

func TestUpdateUser_UsernameShouldBeInvalid(t *testing.T) {
	query := `
        mutation UpdateUserMutation($input: UpdateUserInput!) {
          updateUser(input: $input) {
			user {
				displayName
				picture
			}
			clientMutationId
		  }
        }
	  `
	params := map[string]interface{}{
		"input": map[string]interface{}{
			"username":         "John doe",
			"displayName":      "John doe",
			"picture":          "",
			"clientMutationId": "abcde",
		},
	}

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"updateUser": nil,
		},
	}

	result := graphql.Do(graphql.Params{
		Schema:         Schema,
		RequestString:  query,
		Context:        tgo.John,
		VariableValues: params,
	})

	assert.Equal(t, expected.Data, result.Data, "User should be nil")

	if assert.NotNil(t, result.Errors) {
		assert.Equal(t, "The username field must be alphanum", result.Errors[0].Message, "should not update invalid username")
	}

	query = `
        mutation UpdateUserMutation($input: UpdateUserInput!) {
          updateUser(input: $input) {
			user {
				displayName
				picture
			}
			clientMutationId
		  }
        }
	  `
	params = map[string]interface{}{
		"input": map[string]interface{}{
			"username":         "sn",
			"displayName":      "John doe",
			"picture":          "",
			"clientMutationId": "abcde",
		},
	}

	expected = &graphql.Result{
		Data: map[string]interface{}{
			"updateUser": nil,
		},
	}

	result = graphql.Do(graphql.Params{
		Schema:         Schema,
		RequestString:  query,
		Context:        tgo.John,
		VariableValues: params,
	})

	assert.Equal(t, expected.Data, result.Data, "User should be nil")

	if assert.NotNil(t, result.Errors) {
		assert.Equal(t, "The username field must be between 3 and 50 characters long", result.Errors[0].Message, "should not update invalid username")
	}

	query = `
        mutation UpdateUserMutation($input: UpdateUserInput!) {
          updateUser(input: $input) {
			user {
				displayName
				picture
			}
			clientMutationId
		  }
        }
	  `
	params = map[string]interface{}{
		"input": map[string]interface{}{
			"username":         "snargregrgrrgreg545rgergrgergergergdsvdsvsdvsdvsdvsdv",
			"displayName":      "John doe",
			"picture":          "",
			"clientMutationId": "abcde",
		},
	}

	expected = &graphql.Result{
		Data: map[string]interface{}{
			"updateUser": nil,
		},
	}

	result = graphql.Do(graphql.Params{
		Schema:         Schema,
		RequestString:  query,
		Context:        tgo.John,
		VariableValues: params,
	})

	assert.Equal(t, expected.Data, result.Data, "User should be nil")

	if assert.NotNil(t, result.Errors) {
		assert.Equal(t, "The username field must be between 3 and 50 characters long", result.Errors[0].Message, "should not update invalid username")
	}
}

func TestUpdateUser_ShouldNotBeAbleToUpdateUsername(t *testing.T) {
	query := `
        mutation UpdateUserMutation($input: UpdateUserInput!) {
          updateUser(input: $input) {
			user {
				displayName
				picture
			}
			clientMutationId
		  }
        }
	  `
	params := map[string]interface{}{
		"input": map[string]interface{}{
			"username":         "john",
			"displayName":      "John doe",
			"picture":          "",
			"clientMutationId": "abcde",
		},
	}

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"updateUser": nil,
		},
	}

	result := graphql.Do(graphql.Params{
		Schema:         Schema,
		RequestString:  query,
		Context:        tgo.Albert,
		VariableValues: params,
	})

	assert.Equal(t, expected.Data, result.Data, "User should be nil")

	if assert.NotNil(t, result.Errors) {
		assert.Equal(t, "errors.user.usernameAlreadyDefined", result.Errors[0].Message, "should not update invalid username")
	}
}

func TestUpdateUser_ShouldUpdateUser(t *testing.T) {
	query := `
        mutation UpdateUserMutation($input: UpdateUserInput!) {
          updateUser(input: $input) {
			user {
				username
				displayName
				picture
			}
		  }
        }
	  `
	params := map[string]interface{}{
		"input": map[string]interface{}{
			"username":         "john",
			"displayName":      "John doe",
			"picture":          "",
			"clientMutationId": "abcde",
		},
	}

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"updateUser": map[string]interface{}{
				"user": map[string]interface{}{
					"username":    "john",
					"displayName": "John doe",
					"picture":     nil,
				},
			},
		},
	}

	result := graphql.Do(graphql.Params{
		Schema:         Schema,
		RequestString:  query,
		Context:        tgo.John,
		VariableValues: params,
	})

	assert.Equal(t, expected, result, "Should update user with empty picture")

	query = `
        mutation UpdateUserMutation($input: UpdateUserInput!) {
          updateUser(input: $input) {
			user {
				displayName
				picture
			}
		  }
        }
	  `
	params = map[string]interface{}{
		"input": map[string]interface{}{
			"displayName":      "John doe",
			"picture":          "https://c1.staticflickr.com/4/3310/4568602271_2dbaf43615_b.jpg",
			"clientMutationId": "abcde",
		},
	}

	expected = &graphql.Result{
		Data: map[string]interface{}{
			"updateUser": map[string]interface{}{
				"user": map[string]interface{}{
					"displayName": "John doe",
					"picture":     "https://c1.staticflickr.com/4/3310/4568602271_2dbaf43615_b.jpg",
				},
			},
		},
	}

	result = graphql.Do(graphql.Params{
		Schema:         Schema,
		RequestString:  query,
		Context:        tgo.John,
		VariableValues: params,
	})

	assert.Equal(t, expected, result, "Should update user with picture")
}

func TestChangePassword_AnonymousAccessShouldBeDenied(t *testing.T) {
	query := `
        mutation ChangePasswordMutation($input: ChangePasswordInput!) {
          changePassword(input: $input) {
			user {
				hasPassword
			}
			clientMutationId
		  }
        }
	  `
	params := map[string]interface{}{
		"input": map[string]interface{}{
			"password":         "Anonymous",
			"currentPassword":  "Anonymous",
			"clientMutationId": "abcde",
		},
	}

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"changePassword": nil,
		},
	}

	result := graphql.Do(graphql.Params{
		Schema:         Schema,
		RequestString:  query,
		Context:        tgo.Anonymous,
		VariableValues: params,
	})

	assert.Equal(t, expected.Data, result.Data, "User should be nil")

	if assert.NotNil(t, result.Errors) {
		assert.Equal(t, "Anonymous access is denied", result.Errors[0].Message, "Anonymous should not change password")
	}
}

func TestChangePassword_ShouldNotUpdateInvalidPassword(t *testing.T) {
	query := `
        mutation ChangePasswordMutation($input: ChangePasswordInput!) {
          changePassword(input: $input) {
			user {
				hasPassword
			}
			clientMutationId
		  }
        }
	  `
	params := map[string]interface{}{
		"input": map[string]interface{}{
			"password":         "Inv",
			"currentPassword":  "azertyuiop",
			"clientMutationId": "abcde",
		},
	}

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"changePassword": nil,
		},
	}

	result := graphql.Do(graphql.Params{
		Schema:         Schema,
		RequestString:  query,
		Context:        tgo.John,
		VariableValues: params,
	})

	assert.Equal(t, expected.Data, result.Data, "User should be nil")

	if assert.NotNil(t, result.Errors) {
		assert.Equal(t, "The password field must be between 8 and 20 characters long", result.Errors[0].Message, "should not update invalid password")
	}

	query = `
        mutation ChangePasswordMutation($input: ChangePasswordInput!) {
          changePassword(input: $input) {
			user {
				hasPassword
			}
			clientMutationId
		  }
        }
	  `
	params = map[string]interface{}{
		"input": map[string]interface{}{
			"password":         "azertyuiopazertyuiopa",
			"currentPassword":  "azertyuiop",
			"clientMutationId": "abcde",
		},
	}

	expected = &graphql.Result{
		Data: map[string]interface{}{
			"changePassword": nil,
		},
	}

	result = graphql.Do(graphql.Params{
		Schema:         Schema,
		RequestString:  query,
		Context:        tgo.John,
		VariableValues: params,
	})

	assert.Equal(t, expected.Data, result.Data, "User should be nil")

	if assert.NotNil(t, result.Errors) {
		assert.Equal(t, "The password field must be between 8 and 20 characters long", result.Errors[0].Message, "should not update invalid password")
	}
}

func TestChangePassword_ShouldNotUpdateBadPassword(t *testing.T) {
	query := `
        mutation ChangePasswordMutation($input: ChangePasswordInput!) {
          changePassword(input: $input) {
			user {
				hasPassword
			}
			clientMutationId
		  }
        }
	  `
	params := map[string]interface{}{
		"input": map[string]interface{}{
			"password":         "newpassword",
			"currentPassword":  "badpassword",
			"clientMutationId": "abcde",
		},
	}

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"changePassword": nil,
		},
	}

	result := graphql.Do(graphql.Params{
		Schema:         Schema,
		RequestString:  query,
		Context:        tgo.Albert,
		VariableValues: params,
	})

	assert.Equal(t, expected.Data, result.Data, "User should be nil")

	if assert.NotNil(t, result.Errors) {
		assert.Equal(t, "errors.user.badPassword", result.Errors[0].Message, "Anonymous should not change password")
	}
}

func TestChangePassword_ShouldUpdatePasswordWithoutCurrent(t *testing.T) {
	johnM, _ := models.Users(qm.Where("id = ?", "4c9f32df-c112-4d02-a459-3493fac49ea9")).OneG()
	assert.Nil(t, johnM.Password.Ptr())
	tokenVersion := *johnM.TokenVersion.Ptr()

	query := `
        mutation ChangePasswordMutation($input: ChangePasswordInput!) {
          changePassword(input: $input) {
			user {
				hasPassword
			}
			clientMutationId
		  }
        }
	  `
	params := map[string]interface{}{
		"input": map[string]interface{}{
			"password":         "azertyuiop",
			"currentPassword":  "",
			"clientMutationId": "abcde",
		},
	}

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"changePassword": map[string]interface{}{
				"user": map[string]interface{}{
					"hasPassword": true,
				},
				"clientMutationId": "abcde",
			},
		},
	}

	result := graphql.Do(graphql.Params{
		Schema:         Schema,
		RequestString:  query,
		Context:        tgo.John,
		VariableValues: params,
	})

	johnM, _ = models.Users(qm.Where("id = ?", "4c9f32df-c112-4d02-a459-3493fac49ea9")).OneG()

	assert.Equal(t, expected, result, "Should update password without current password if has no password")
	assert.NotNil(t, johnM.Password.Ptr(), "Should not be null")
	assert.NotEqual(t, tokenVersion, *johnM.TokenVersion.Ptr())
}

func TestChangePassword_ShouldUpdatePassword(t *testing.T) {
	query := `
        mutation ChangePasswordMutation($input: ChangePasswordInput!) {
          changePassword(input: $input) {
			user {
				hasPassword
			}
			clientMutationId
		  }
        }
	  `
	params := map[string]interface{}{
		"input": map[string]interface{}{
			"password":         "azertyuiopoiu",
			"currentPassword":  "azertyuiop",
			"clientMutationId": "abcde",
		},
	}

	expected := &graphql.Result{
		Data: map[string]interface{}{
			"changePassword": map[string]interface{}{
				"user": map[string]interface{}{
					"hasPassword": true,
				},
				"clientMutationId": "abcde",
			},
		},
	}

	result := graphql.Do(graphql.Params{
		Schema:         Schema,
		RequestString:  query,
		Context:        tgo.John,
		VariableValues: params,
	})

	assert.Equal(t, expected, result, "Should update password")
}
