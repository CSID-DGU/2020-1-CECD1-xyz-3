import {
  ChangeDetectionStrategy,
  Component,
} from '@angular/core';

@Component({
  selector: 'dg-admin-preference',
  templateUrl: './preference.component.html',
  styleUrls: [ './preference.component.scss' ],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class PreferenceComponent {

}
