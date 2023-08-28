import { Context } from 'aws-lambda';
import { Logger } from '@aws-lambda-powertools/logger';
import {handlerLogic} from "./src/handler";

{%- if cookiecutter["Powertools for AWS Lambda (TypeScript) Metrics"] == "enabled"%}
import { Metrics, MetricUnits } from '@aws-lambda-powertools/metrics';
{%- endif %}
{%- if cookiecutter["Powertools for AWS Lambda (TypeScript) Tracing"] == "enabled"%}
import { Tracer } from '@aws-lambda-powertools/tracer';
{%- endif %}


{%- if cookiecutter["Powertools for AWS Lambda (TypeScript) Metrics"] == "enabled"%}
const metrics = new Metrics();
{%- endif %}
const logger = new Logger();
{%- if cookiecutter["Powertools for AWS Lambda (TypeScript) Tracing"] == "enabled"%}
const tracer = new Tracer();
{%- endif %}

export const lambdaHandler = async (event: any, context: Context): Promise<any> => {
    logger.info('Lambda invocation event', { event });
    logger.appendKeys({
        awsRequestId: context.awsRequestId,
    });

    {%- if cookiecutter["Powertools for AWS Lambda (TypeScript) Tracing"] == "enabled"%}
    // Get facade segment created by AWS Lambda
    const segment = tracer.getSegment();

    if (!segment) {
        throw new Error('Failed to get segment');
    }

    // Create subsegment for the function & set it as active
    const handlerSegment = segment.addNewSubsegment(`## ${process.env._HANDLER}`);
    tracer.setSegment(handlerSegment);

    // Annotate the subsegment with the cold start & serviceName
    tracer.annotateColdStart();
    tracer.addServiceNameAnnotation();

    // Add annotation for the awsRequestId
    tracer.putAnnotation('awsRequestId', context.awsRequestId);

    {%- endif %}
    {%- if cookiecutter["Powertools for AWS Lambda (TypeScript) Metrics"] == "enabled" %}
    // Capture cold start metrics
    metrics.captureColdStartMetric();

    {%- endif %}
    {%- if cookiecutter["Powertools for AWS Lambda (TypeScript) Tracing"] == "enabled"%}
    // Create another subsegment & set it as active
    const subsegment = handlerSegment.addNewSubsegment('### MySubSegment');
    tracer.setSegment(subsegment);
    {%- endif %}

    try {
        return  handlerLogic(event)
    } catch (err) {
        {%- if cookiecutter["Powertools for AWS Lambda (TypeScript) Tracing"] == "enabled"%}
        tracer.addErrorAsMetadata(err as Error);
        {%- endif %}

        logger.error(`Error occurred: ${err}`, event);
    {%- if cookiecutter["Powertools for AWS Lambda (TypeScript) Metrics"] == "enabled" or cookiecutter["Powertools for AWS Lambda (TypeScript) Tracing"] == "enabled"%}
    } finally {
        {%- if cookiecutter["Powertools for AWS Lambda (TypeScript) Tracing"] == "enabled"%}
        // Close subsegments (the AWS Lambda one is closed automatically)
        subsegment.close(); // (### MySubSegment)
        handlerSegment.close(); // (## index.handler)

        // Set the facade segment as active again (the one created by AWS Lambda)
        tracer.setSegment(segment);

        {%- endif %}
        {%- if cookiecutter["Powertools for AWS Lambda (TypeScript) Metrics"] == "enabled"%}
        // Publish all stored metrics
        metrics.publishStoredMetrics();

        {%- endif %}
    {%- endif %}
    }

};