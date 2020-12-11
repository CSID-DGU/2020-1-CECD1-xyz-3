import {
  ChangeDetectionStrategy,
  Component,
} from '@angular/core';
import { Router } from '@angular/router';
import { ApiService } from 'app/core';
import { TranslateService } from 'app/core/translate.service';

interface AdminMenu {
  label: string;
  link: string;
}

@Component({
  selector: 'dg-admin',
  templateUrl: './admin.component.html',
  styleUrls: [ './admin.component.scss' ],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AdminComponent {
  public readonly menus: AdminMenu[] = [
    // { label: this.translateService.translate('DASHBOARD'), link: 'dashboard' },
    { label: this.translateService.translate('APPLICATION_MANAGEMENT'), link: 'applicant-management' },
    { label: this.translateService.translate('POSITION_MANAGEMENT'), link: 'position-management' },
    { label: this.translateService.translate('FORM_MANAGEMENT'), link: 'form-management' },
  ];

  public constructor(
    private readonly apiService: ApiService,
    private readonly router: Router,
    private readonly translateService: TranslateService,
  ) {
    return;
  }

  public isActiveRouter(link: string): boolean {
    return this.router.url.lastIndexOf(link) === this.router.url.length - link.length;
  }

  public logout(): void {
    this.apiService.logout()
      .subscribe(() => {
        this.router.navigate([ `/login` ], { queryParams: { redirectTo: window.location.pathname } });
      }, () => {
        alert(this.translateService.translate('ERROR_LOGOUT_FAILED'));
      });
  }
}
