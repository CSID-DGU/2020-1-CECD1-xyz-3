import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { JobListModule } from './job-list/job-list.module';
import { LoginGuard } from './login/login-guard';
import { CoreModule } from './core/core.module';
import { LoginComponent } from './login/login.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    JobListModule,
    CoreModule,
    CommonModule,
    FormsModule,
  ],
  providers: [
    LoginGuard,
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
