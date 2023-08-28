import {Context} from 'aws-lambda';
import {lambdaHandler} from '../../app';
import {describe, expect, it} from '@jest/globals';

describe('Unit test for app handler', function () {
    it('verifies successful response', async () => {
        const event = {

        };
        const context: Context = {
            callbackWaitsForEmptyEventLoop: false,
            functionName: 'lambdaHandler',
            functionVersion: '1.0',
            invokedFunctionArn: 'arn:1234567890:lambda:lambdaHandler',
            memoryLimitInMB: '128',
            awsRequestId: '1234567890',
            logGroupName: 'lambdaHandlerLogGroup',
            logStreamName: 'c6a789dff9326bc178',
            getRemainingTimeInMillis: function (): number {
                throw new Error('Function not implemented.');
            },
            done: function (error?: Error, result?: any): void {
                throw new Error('Function not implemented.');
            },
            fail: function (error: string | Error): void {
                throw new Error('Function not implemented.');
            },
            succeed: function (messageOrObject: any): void {
                throw new Error('Function not implemented.');
            }
        };
        const result: any = await lambdaHandler(event,context);

        expect(result).toEqual({
            message: 'hello world',
        });
    });
});
