import { nestHttpExceptionFilter } from './nest-http-exception-filter';

describe('nestHttpExceptionFilter', () => {
  it('should work', () => {
    expect(nestHttpExceptionFilter()).toEqual('nest-http-exception-filter');
  });
});
