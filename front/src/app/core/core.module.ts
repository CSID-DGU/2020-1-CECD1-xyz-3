import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { ApiService } from './api.service';
import { TranslateService } from './translate.service';
import { TranslatePipe } from './translate.pipe';

@NgModule({
  declarations: [
    TranslatePipe,
  ],
  imports: [
    HttpClientModule,
  ],
  exports: [
    TranslatePipe,
  ],
  providers: [
    ApiService,
    TranslateService,
  ],
})
export class CoreModule {}
