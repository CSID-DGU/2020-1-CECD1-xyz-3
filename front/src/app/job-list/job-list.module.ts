import { NgModule, InjectionToken } from '@angular/core';
import { JobListComponent } from './job-list.component';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { CoreModule } from 'app/core/core.module';

@NgModule({
  declarations: [
    JobListComponent,
  ],
  imports: [
    CommonModule,
    RouterModule,
    FormsModule,
    CoreModule,
  ],
  exports: [],
  providers: [
    { provide: 'userName', useValue: { name: undefined, email: undefined } },
  ],
})
export class JobListModule {}
