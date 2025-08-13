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
pub enum GetProtectedResourceResponse {
    /// Congrats, you captured the protected flag.
    Status200_Congrats
    (models::ProtectedResource)
    ,
    /// The request is unauthorized
    Status401_TheRequestIsUnauthorized
    (models::ErrorMessage)
    ,
    /// Unexpected server error
    Status500_UnexpectedServerError
    (models::ErrorMessage)
}


/// Protected
#[async_trait]
#[allow(clippy::ptr_arg)]
pub trait Protected<E: std::fmt::Debug + Send + Sync + 'static = ()>: super::ErrorHandler<E> {
    type Claims;

    /// Get a protected resource.
    ///
    /// GetProtectedResource - GET /protected
    async fn get_protected_resource(
    &self,
    method: &Method,
    host: &Host,
    cookies: &CookieJar,
        claims: &Self::Claims,
    ) -> Result<GetProtectedResourceResponse, E>;
}
