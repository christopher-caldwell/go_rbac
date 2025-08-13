use async_trait::async_trait;
use axum::extract::*;
use axum_extra::extract::{CookieJar, Host};
use bytes::Bytes;
use http::Method;
use serde::{Deserialize, Serialize};

use crate::{models, types::*};

#[derive(Debug, PartialEq, Serialize, Deserialize)]
#[must_use]
#[allow(clippy::large_enum_variant)]
pub enum CreateRbacResourceResponse {
    /// Congrats, you captured the rbac flag.
    Status201_Congrats
    (models::RbacResource)
    ,
    /// Bad input
    Status400_BadInput
    (models::ErrorMessage)
    ,
    /// The request is unauthorized
    Status401_TheRequestIsUnauthorized
    (models::ErrorMessage)
    ,
    /// Unexpected server error
    Status500_UnexpectedServerError
    (models::ErrorMessage)
}

#[derive(Debug, PartialEq, Serialize, Deserialize)]
#[must_use]
#[allow(clippy::large_enum_variant)]
pub enum GetRbacResourceResponse {
    /// Congrats, you captured the rbac flag.
    Status200_Congrats
    (models::RbacResource)
    ,
    /// The request is unauthorized
    Status401_TheRequestIsUnauthorized
    (models::ErrorMessage)
    ,
    /// The specified resource was not found
    Status404_TheSpecifiedResourceWasNotFound
    (models::ErrorMessage)
    ,
    /// Unexpected server error
    Status500_UnexpectedServerError
    (models::ErrorMessage)
}


/// Rbac
#[async_trait]
#[allow(clippy::ptr_arg)]
pub trait Rbac<E: std::fmt::Debug + Send + Sync + 'static = ()>: super::ErrorHandler<E> {
    type Claims;

    /// Create a resource that requires rbac.
    ///
    /// CreateRbacResource - POST /rbac
    async fn create_rbac_resource(
    &self,
    method: &Method,
    host: &Host,
    cookies: &CookieJar,
        claims: &Self::Claims,
            body: &models::RbacResource,
    ) -> Result<CreateRbacResourceResponse, E>;

    /// Get a resource that requires rbac.
    ///
    /// GetRbacResource - GET /rbac
    async fn get_rbac_resource(
    &self,
    method: &Method,
    host: &Host,
    cookies: &CookieJar,
        claims: &Self::Claims,
    ) -> Result<GetRbacResourceResponse, E>;
}
