import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { BackendDocsModule } from '@nx-multi-cloud/nest/swagger-middleware';

@Module({
  imports: [BackendDocsModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
