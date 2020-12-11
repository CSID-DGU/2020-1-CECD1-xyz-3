import { NgModule } from '@angular/core';
import {
  RouterModule,
  Routes,
} from '@angular/router';
import { AdminDashboardComponent } from './dashboard/dashboard.component';
import { AdminComponent } from './admin.component';
import { ApplicantManagementComponent } from './applicant-management/applicant-management.component';
import { PositionManagementComponent } from './position-management/position-management.component';
import { FormManagementComponent } from './form-management/form-management.component';
import { PreferenceComponent } from './preference/preference.component';


const routes: Routes = [
  {
    path: '',
    component: AdminComponent,
    children: [
      { path: '', pathMatch: 'full', redirectTo: 'applicant-management' },
      { path: 'dashboard', component: AdminDashboardComponent },
      { path: 'applicant-management', component: ApplicantManagementComponent },
      { path: 'position-management', component: PositionManagementComponent },
      { path: 'form-management', component: FormManagementComponent },
      { path: 'preference', component: PreferenceComponent },
    ],
  },
];

@NgModule({
  imports: [ RouterModule.forChild(routes) ],
  exports: [ RouterModule ],
})
export class AdminRoutingModule {}
