import {
  ChangeDetectionStrategy,
  Component,
} from '@angular/core';

@Component({
  selector: 'dg-admin-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: [ './dashboard.component.scss' ],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AdminDashboardComponent {

}
