import { nestSwaggerMiddleware } from './nest-swagger-middleware';

describe('nestSwaggerMiddleware', () => {
  it('should work', () => {
    expect(nestSwaggerMiddleware()).toEqual('nest-swagger-middleware');
  });
});
