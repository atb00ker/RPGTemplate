import React from 'react';
import { ErrorBoundary } from 'react-error-boundary';
import { ErrorFallback } from '../ErrorFallback/ErrorFallback';
import { Link } from "react-router-dom";

const RPGTemplate = () => {
  return (
    <React.Fragment>
      <ErrorBoundary FallbackComponent={ErrorFallback}>
      </ErrorBoundary>
    </React.Fragment>
  );
};

export default RPGTemplate;
