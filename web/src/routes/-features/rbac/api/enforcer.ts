import * as casbinjs from 'casbin.js'

// Set your backend Casbin service URL
export const authorizer = new casbinjs.Authorizer('manual')
