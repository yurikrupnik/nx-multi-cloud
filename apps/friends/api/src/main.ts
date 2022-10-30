// import {
//   FastifyAdapter,
//   NestFastifyApplication,
// } from '@nestjs/platform-fastify';
import { Logger } from '@nestjs/common';
import helmet from 'helmet';
// import helmet from '@fastify/helmet';
import { NestFactory } from '@nestjs/core';
import { AppModule } from './app/app.module';
import { BackendDocsModule } from '@nx-multi-cloud/nest/swagger-middleware';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  const globalPrefix = 'api';
  app.setGlobalPrefix(globalPrefix);

  app.use(helmet());
  app.use(helmet.noSniff());
  app.use(helmet.hidePoweredBy());
  app.use(helmet.contentSecurityPolicy());

  app.setGlobalPrefix(globalPrefix);
  app.enableShutdownHooks();

  const docs = app.get(BackendDocsModule);
  docs.setup(app, globalPrefix, 'Friends API', 'General use cloud run api');

  const port = process.env.PORT || 8080;
  await app.listen(port, '0.0.0.0');
  Logger.log(
    `🚀 Application is running on: http://localhost:${port}/${globalPrefix}`
  );
}

bootstrap();
