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
pub enum GetUnprotectedResourceResponse {
    /// Congrats, you captured the unprotected flag.
    Status200_Congrats
    (models::UnprotectedResource)
    ,
    /// Unexpected server error
    Status500_UnexpectedServerError
    (models::ErrorMessage)
}


/// Unprotected
#[async_trait]
#[allow(clippy::ptr_arg)]
pub trait Unprotected<E: std::fmt::Debug + Send + Sync + 'static = ()>: super::ErrorHandler<E> {
    /// Get an unprotected resource.
    ///
    /// GetUnprotectedResource - GET /unprotected
    async fn get_unprotected_resource(
    &self,
    method: &Method,
    host: &Host,
    cookies: &CookieJar,
    ) -> Result<GetUnprotectedResourceResponse, E>;
}
