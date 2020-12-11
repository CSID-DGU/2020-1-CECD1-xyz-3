import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { AdminRoutingModule } from './admin-routing.module';
import { AdminComponent } from './admin.component';
import { AdminDashboardComponent } from './dashboard/dashboard.component';
import { ApplicantManagementComponent } from './applicant-management/applicant-management.component';
import { PositionManagementComponent } from './position-management/position-management.component';
import { CoreModule } from 'app/core/core.module';

@NgModule({
  declarations: [
    AdminComponent,
    AdminDashboardComponent,
    ApplicantManagementComponent,
    PositionManagementComponent,
  ],
  imports: [
    AdminRoutingModule,
    CommonModule,
    FormsModule,
    CoreModule,
  ],
  exports: [],
  providers: [],
})
export class AdminModule {}
