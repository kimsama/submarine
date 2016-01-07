﻿// This file was generated by typhen-api

using System;

namespace TyphenApi
{
    public class WebSocketSessionError<T> : SystemException where T : TypeBase
    {
        public T Error { get; private set; }
        public Exception RawError { get; private set; }
        public string RawErrorMessage { get; private set; }

        public WebSocketSessionError(T error, Exception rawError, string rawErrorMessage)
            : base(rawErrorMessage)
        {
            Error = error;
            RawError = rawError;
            RawErrorMessage = rawErrorMessage;
        }

        public override string ToString()
        {
            return string.Format(
                "[WebSocketSessionError: RawError={0}, RawErrorMessage={1}, Error={2}]",
                RawError, RawErrorMessage, Error);
        }
    }
}
