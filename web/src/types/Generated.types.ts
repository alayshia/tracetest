/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */

export interface paths {
  "/tests": {
    /** get tests */
    get: operations["getTests"];
    /** Create new test action */
    post: operations["createTest"];
  };
  "/tests/definition.yaml": {
    /** Create new test using the yaml definition */
    post: operations["createTestFromDefinition"];
  };
  "/tests/{testId}/definition.yaml": {
    /** update test action from definition file */
    put: operations["updateTestFromDefinition"];
  };
  "/tests/{testId}": {
    /** get test */
    get: operations["getTest"];
    /** update test action */
    put: operations["updateTest"];
    /** delete a test */
    delete: operations["deleteTest"];
  };
  "/tests/{testId}/run": {
    /** get the runs from a particular test */
    get: operations["getTestRuns"];
    /** run a particular test */
    post: operations["runTest"];
  };
  "/tests/{testId}/run/{runId}/select": {
    /** get the spans ids that would be selected by a specific selector query */
    get: operations["getTestResultSelectedSpans"];
  };
  "/tests/{testId}/run/{runId}/dry-run": {
    /** use this method to test a definition against an actual trace without creating a new version or persisting anything */
    put: operations["dryRunAssertion"];
  };
  "/tests/{testId}/run/{runId}/rerun": {
    /** rerun a test run */
    post: operations["rerunTestRun"];
  };
  "/tests/{testId}/run/{runId}/junit.xml": {
    /** get test run results in JUnit xml format */
    get: operations["getRunResultJUnit"];
  };
  "/tests/{testId}/run/{runId}/export": {
    /** export test and test run information for debugging */
    get: operations["exportTestRun"];
  };
  "/tests/import": {
    /** import test and test run information for debugging */
    post: operations["importTestRun"];
  };
  "/tests/{testId}/run/{runId}": {
    /** get a particular test Run */
    get: operations["getTestRun"];
    /** delete a test run */
    delete: operations["deleteTestRun"];
  };
  "/tests/{testId}/definition": {
    /** Gets definition for a test */
    get: operations["getTestSpecs"];
    /** Set spec for a particular test */
    put: operations["setTestSpecs"];
  };
  "/tests/{testId}/version/{version}": {
    /** get a test specific version */
    get: operations["getTestVersion"];
  };
  "/tests/{testId}/version/{version}/definition.yaml": {
    /** Get the test definition as an YAML file */
    get: operations["getTestVersionDefinitionFile"];
  };
}

export interface components {}

export interface operations {
  /** get tests */
  getTests: {
    parameters: {
      query: {
        /** indicates how many tests can be returned by each page */
        take?: number;
        /** indicates how many tests will be skipped when paginating */
        skip?: number;
        /** query to search tests, based on test name and description */
        query?: string;
        /** indicates the sort field for the tests */
        sortBy?: "created" | "name" | "last_run";
        /** indicates the sort direction for the tests */
        sortDirection?: "asc" | "desc";
      };
    };
    responses: {
      /** successful operation */
      200: {
        headers: {
          /** Total records count */
          "X-Total-Count"?: number;
        };
        content: {
          "application/json": external["tests.yaml"]["components"]["schemas"]["Test"][];
        };
      };
      /** problem with getting tests */
      500: unknown;
    };
  };
  /** Create new test action */
  createTest: {
    responses: {
      /** successful operation */
      200: {
        content: {
          "application/json": external["tests.yaml"]["components"]["schemas"]["Test"];
        };
      };
      /** trying to create a test with an already existing ID */
      400: unknown;
    };
    requestBody: {
      content: {
        "application/json": external["tests.yaml"]["components"]["schemas"]["Test"];
      };
    };
  };
  /** Create new test using the yaml definition */
  createTestFromDefinition: {
    responses: {
      /** successful operation */
      200: {
        content: {
          "application/json": external["tests.yaml"]["components"]["schemas"]["Test"];
        };
      };
      /** trying to create a test with an already existing ID */
      400: unknown;
    };
    requestBody: {
      content: {
        "text/json": external["tests.yaml"]["components"]["schemas"]["TextDefinition"];
      };
    };
  };
  /** update test action from definition file */
  updateTestFromDefinition: {
    parameters: {
      path: {
        testId: string;
      };
    };
    responses: {
      /** successful operation */
      200: {
        content: {
          "application/json": external["tests.yaml"]["components"]["schemas"]["Test"];
        };
      };
    };
    requestBody: {
      content: {
        "text/json": external["tests.yaml"]["components"]["schemas"]["TextDefinition"];
      };
    };
  };
  /** get test */
  getTest: {
    parameters: {
      path: {
        testId: string;
      };
    };
    responses: {
      /** successful operation */
      200: {
        content: {
          "application/json": external["tests.yaml"]["components"]["schemas"]["Test"];
        };
      };
      /** problem with getting a test */
      500: unknown;
    };
  };
  /** update test action */
  updateTest: {
    parameters: {
      path: {
        testId: string;
      };
    };
    responses: {
      /** successful operation */
      204: never;
      /** problem with updating test */
      500: unknown;
    };
    requestBody: {
      content: {
        "application/json": external["tests.yaml"]["components"]["schemas"]["Test"];
      };
    };
  };
  /** delete a test */
  deleteTest: {
    parameters: {
      path: {
        testId: string;
      };
    };
    responses: {
      /** OK */
      204: never;
    };
  };
  /** get the runs from a particular test */
  getTestRuns: {
    parameters: {
      path: {
        testId: string;
      };
      query: {
        /** indicates how many results can be returned by each page */
        take?: number;
        /** indicates how many results will be skipped when paginating */
        skip?: number;
      };
    };
    responses: {
      /** successful operation */
      200: {
        headers: {
          /** Total records count */
          "X-Total-Count"?: number;
        };
        content: {
          "application/json": external["tests.yaml"]["components"]["schemas"]["TestRun"][];
        };
      };
    };
  };
  /** run a particular test */
  runTest: {
    parameters: {
      path: {
        testId: string;
      };
    };
    responses: {
      /** successful operation */
      200: {
        content: {
          "application/json": external["tests.yaml"]["components"]["schemas"]["TestRun"];
        };
      };
    };
    requestBody: {
      content: {
        "application/json": external["tests.yaml"]["components"]["schemas"]["TestRunInformation"];
      };
    };
  };
  /** get the spans ids that would be selected by a specific selector query */
  getTestResultSelectedSpans: {
    parameters: {
      path: {
        testId: string;
        runId: string;
      };
      query: {
        query?: string;
      };
    };
    responses: {
      /** successful operation */
      200: {
        content: {
          "application/json": external["tests.yaml"]["components"]["schemas"]["SelectedSpansResult"];
        };
      };
    };
  };
  /** use this method to test a definition against an actual trace without creating a new version or persisting anything */
  dryRunAssertion: {
    parameters: {
      path: {
        testId: string;
        runId: string;
      };
    };
    responses: {
      /** successful operation */
      200: {
        content: {
          "application/json": external["tests.yaml"]["components"]["schemas"]["AssertionResults"];
        };
      };
    };
    requestBody: {
      content: {
        "application/json": external["tests.yaml"]["components"]["schemas"]["TestSpecs"];
      };
    };
  };
  /** rerun a test run */
  rerunTestRun: {
    parameters: {
      path: {
        testId: string;
        runId: string;
      };
    };
    responses: {
      /** successful operation */
      200: {
        content: {
          "application/json": external["tests.yaml"]["components"]["schemas"]["TestRun"];
        };
      };
    };
  };
  /** get test run results in JUnit xml format */
  getRunResultJUnit: {
    parameters: {
      path: {
        testId: string;
        runId: string;
      };
    };
    responses: {
      /** JUnit formatted file */
      200: {
        content: {
          "application/xml": string;
        };
      };
    };
  };
  /** export test and test run information for debugging */
  exportTestRun: {
    parameters: {
      path: {
        testId: string;
        runId: string;
      };
    };
    responses: {
      /** successfuly exported test and test run information */
      200: {
        content: {
          "application/json": external["tests.yaml"]["components"]["schemas"]["ExportedTestInformation"];
        };
      };
    };
  };
  /** import test and test run information for debugging */
  importTestRun: {
    responses: {
      /** successfuly imported test and test run information */
      200: {
        content: {
          "application/json": external["tests.yaml"]["components"]["schemas"]["ExportedTestInformation"];
        };
      };
    };
    requestBody: {
      content: {
        "application/json": external["tests.yaml"]["components"]["schemas"]["ExportedTestInformation"];
      };
    };
  };
  /** get a particular test Run */
  getTestRun: {
    parameters: {
      path: {
        testId: string;
        runId: string;
      };
    };
    responses: {
      /** successful operation */
      200: {
        content: {
          "application/json": external["tests.yaml"]["components"]["schemas"]["TestRun"];
        };
      };
    };
  };
  /** delete a test run */
  deleteTestRun: {
    parameters: {
      path: {
        testId: string;
        runId: string;
      };
    };
    responses: {
      /** OK */
      204: never;
    };
  };
  /** Gets definition for a test */
  getTestSpecs: {
    parameters: {
      path: {
        testId: string;
      };
    };
    responses: {
      /** successful operation */
      201: {
        content: {
          "application/json": external["tests.yaml"]["components"]["schemas"]["TestSpecs"][];
        };
      };
    };
  };
  /** Set spec for a particular test */
  setTestSpecs: {
    parameters: {
      path: {
        testId: string;
      };
    };
    responses: {
      /** OK */
      204: never;
    };
    requestBody: {
      content: {
        "application/json": external["tests.yaml"]["components"]["schemas"]["TestSpecs"];
      };
    };
  };
  /** get a test specific version */
  getTestVersion: {
    parameters: {
      path: {
        testId: string;
        version: number;
      };
    };
    responses: {
      /** successful operation */
      200: {
        content: {
          "application/json": external["tests.yaml"]["components"]["schemas"]["Test"];
        };
      };
      /** problem with getting a test */
      500: unknown;
    };
  };
  /** Get the test definition as an YAML file */
  getTestVersionDefinitionFile: {
    parameters: {
      path: {
        testId: string;
        version: number;
      };
    };
    responses: {
      /** OK */
      200: {
        content: {
          "application/yaml": string;
        };
      };
    };
  };
}

export interface external {
  "grpc.yaml": {
    paths: {};
    components: {
      schemas: {
        GRPCHeader: {
          key?: string;
          value?: string;
        };
        GRPCRequest: {
          protobufFile?: string;
          address?: string;
          service?: string;
          method?: string;
          metadata?: external["grpc.yaml"]["components"]["schemas"]["GRPCHeader"][];
          auth?: external["http.yaml"]["components"]["schemas"]["HTTPAuth"];
          request?: string;
        };
        GRPCResponse: {
          statusCode?: number;
          metadata?: external["grpc.yaml"]["components"]["schemas"]["GRPCHeader"][];
          body?: string;
        };
      };
    };
    operations: {};
  };
  "http.yaml": {
    paths: {};
    components: {
      schemas: {
        HTTPHeader: {
          key?: string;
          value?: string;
        };
        HTTPRequest: {
          url?: string;
          /** @enum {string} */
          method?:
            | "GET"
            | "PUT"
            | "POST"
            | "PATCH"
            | "DELETE"
            | "COPY"
            | "HEAD"
            | "OPTIONS"
            | "LINK"
            | "UNLINK"
            | "PURGE"
            | "LOCK"
            | "UNLOCK"
            | "PROPFIND"
            | "VIEW";
          headers?: external["http.yaml"]["components"]["schemas"]["HTTPHeader"][];
          body?: string;
          auth?: external["http.yaml"]["components"]["schemas"]["HTTPAuth"];
        };
        HTTPResponse: {
          status?: string;
          statusCode?: number;
          headers?: external["http.yaml"]["components"]["schemas"]["HTTPHeader"][];
          body?: string;
        };
        HTTPAuth: {
          /** @enum {string} */
          type?: "apiKey" | "basic" | "bearer";
          apiKey?: {
            key?: string;
            value?: string;
            /** @enum {string} */
            in?: "query" | "header";
          };
          basic?: {
            username?: string;
            password?: string;
          };
          bearer?: {
            token?: string;
          };
        };
      };
    };
    operations: {};
  };
  "tests.yaml": {
    paths: {};
    components: {
      schemas: {
        Test: {
          id?: string;
          name?: string;
          description?: string;
          /** @description version number of the test */
          version?: number;
          serviceUnderTest?: external["triggers.yaml"]["components"]["schemas"]["Trigger"];
          /** @description specification of assertions that are going to be made */
          specs?: external["tests.yaml"]["components"]["schemas"]["TestSpecs"];
          /** @description summary of test data */
          summary?: external["tests.yaml"]["components"]["schemas"]["TestSummary"];
        };
        TestSummary: {
          runs?: number;
          lastRun?: {
            /** Format: date-time */
            time?: string;
            passes?: number;
            fails?: number;
          };
        };
        /** @example [object Object] */
        TestSpecs: {
          specs?: {
            name?: string | null;
            selector?: external["tests.yaml"]["components"]["schemas"]["Selector"];
            assertions?: string[];
          }[];
        };
        TestRun: {
          id?: string;
          traceId?: string;
          spanId?: string;
          /** @description Test version used when running this test run */
          testVersion?: number;
          /**
           * @description Current execution state
           * @enum {string}
           */
          state?:
            | "CREATED"
            | "EXECUTING"
            | "AWAITING_TRACE"
            | "AWAITING_TEST_RESULTS"
            | "FINISHED"
            | "FAILED";
          /** @description Details of the cause for the last `FAILED` state */
          lastErrorState?: string;
          /** @description time it took for the test to complete, either success or fail. If the test is still running, it will show the time up to the time of the request */
          executionTime?: number;
          /** Format: date-time */
          createdAt?: string;
          /** Format: date-time */
          serviceTriggeredAt?: string;
          /** Format: date-time */
          serviceTriggerCompletedAt?: string;
          /** Format: date-time */
          obtainedTraceAt?: string;
          /** Format: date-time */
          completedAt?: string;
          triggerResult?: external["triggers.yaml"]["components"]["schemas"]["TriggerResult"];
          trace?: external["trace.yaml"]["components"]["schemas"]["Trace"];
          result?: external["tests.yaml"]["components"]["schemas"]["AssertionResults"];
          metadata?: { [key: string]: string };
        };
        TestRunInformation: {
          metadata?: { [key: string]: string } | null;
        };
        /** @example [object Object] */
        AssertionResults: {
          allPassed?: boolean;
          results?: {
            selector?: external["tests.yaml"]["components"]["schemas"]["Selector"];
            results?: external["tests.yaml"]["components"]["schemas"]["AssertionResult"][];
          }[];
        };
        AssertionResult: {
          assertion?: string;
          allPassed?: boolean;
          spanResults?: external["tests.yaml"]["components"]["schemas"]["AssertionSpanResult"][];
        };
        AssertionSpanResult: {
          spanId?: string;
          observedValue?: string;
          passed?: boolean;
          error?: string;
        };
        DefinitionFile: {
          content?: string;
        };
        SelectedSpansResult: {
          selector?: external["tests.yaml"]["components"]["schemas"]["Selector"];
          spanIds?: string[];
        };
        Selector: {
          query?: string;
          structure?: external["tests.yaml"]["components"]["schemas"]["SpanSelector"][];
        };
        SpanSelector: {
          filters: external["tests.yaml"]["components"]["schemas"]["SelectorFilter"][];
          pseudoClass?: external["tests.yaml"]["components"]["schemas"]["SelectorPseudoClass"];
          childSelector?: external["tests.yaml"]["components"]["schemas"]["SpanSelector"];
        } | null;
        SelectorFilter: {
          property: string;
          operator: string;
          value: string;
        };
        SelectorPseudoClass: {
          name: string;
          argument?: number;
        } | null;
        ExportedTestInformation: {
          test: external["tests.yaml"]["components"]["schemas"]["Test"];
          run: external["tests.yaml"]["components"]["schemas"]["TestRun"];
        };
        TextDefinition: {
          content?: string;
        };
      };
    };
    operations: {};
  };
  "trace.yaml": {
    paths: {};
    components: {
      schemas: {
        Trace: {
          traceId?: string;
          tree?: external["trace.yaml"]["components"]["schemas"]["Span"];
          /** @description falttened version, mapped as spanId -> span{} */
          flat?: {
            [
              key: string
            ]: external["trace.yaml"]["components"]["schemas"]["Span"];
          };
        };
        Span: {
          id?: string;
          parentId?: string;
          name?: string;
          /**
           * Format: int64
           * @description span start time in unix milli format
           * @example 1656701595277
           */
          startTime?: number;
          /**
           * Format: int64
           * @description span end time in unix milli format
           * @example 1656701595800
           */
          endTime?: number;
          /**
           * @description Key-Value of span attributes
           * @example [object Object]
           */
          attributes?: { [key: string]: string };
          children?: external["trace.yaml"]["components"]["schemas"]["Span"][];
        };
      };
    };
    operations: {};
  };
  "triggers.yaml": {
    paths: {};
    components: {
      schemas: {
        Trigger: {
          /** @enum {string} */
          triggerType?: "http" | "grpc";
          triggerSettings?: {
            http?: external["http.yaml"]["components"]["schemas"]["HTTPRequest"];
            grpc?: external["grpc.yaml"]["components"]["schemas"]["GRPCRequest"];
          };
        };
        TriggerResult: {
          /** @enum {string} */
          triggerType?: "http" | "grpc";
          triggerResult?: {
            http?: external["http.yaml"]["components"]["schemas"]["HTTPResponse"];
            grpc?: external["grpc.yaml"]["components"]["schemas"]["GRPCResponse"];
          };
        };
      };
    };
    operations: {};
  };
}
