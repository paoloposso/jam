import { Controller, Get } from '@nestjs/common';
import { stringify } from 'querystring';
import { AppService } from './app.service';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get()
  check(): string {
    return this.appService.check()?.toString();
  }
}
