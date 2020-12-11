import {
  ChangeDetectionStrategy,
  ChangeDetectorRef,
  Component,
} from '@angular/core';
import {
  ActivatedRoute,
  Router,
} from '@angular/router';
import { ApiService } from 'app/core';
import * as RxOp from 'rxjs/operators';
import { TranslateService } from 'app/core/translate.service';

@Component({
  selector: 'dg-login',
  templateUrl: './login.component.html',
  styleUrls: [ './login.component.scss' ],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class LoginComponent {
  public companyName: string;
  public redirectTo: string;
  public userId: string = '';
  public userPassword: string = '';
  public isProcessingLogin: boolean = false;

  private companyDomain: string;

  public constructor(
    private readonly apiService: ApiService,
    private readonly changeDetectorRef: ChangeDetectorRef,
    private readonly route: ActivatedRoute,
    private readonly router: Router,
    private readonly translateService: TranslateService,
  ) {
    this.redirectTo = route.snapshot.queryParamMap.get('redirectTo');
    this.companyDomain = this.redirectTo.replace(/\/([^/]+)\/.*/, '$1');
    if (/^[a-zA-Z0-9]+\.[a-zA-Z0-9]+$/.test(this.companyDomain)) {
      this.companyName = this.companyDomain.substr(0, this.companyDomain.indexOf('.'));
    } else {
      alert(this.translateService.translate('ERROR_INVALID_PATH'));
    }
  }

  public login(): void {
    if (!(new RegExp(`^[^@]+@${this.companyDomain}$`).test(this.userId))) {
      alert(this.translateService.translate('ERROR_INVALID_ID_OR_PASSWORD'));
      return;
    }
    this.isProcessingLogin = true;
    this.apiService.login(this.userId, this.userPassword)
      .pipe(
        RxOp.finalize(() => {
          this.isProcessingLogin = false;
          this.changeDetectorRef.markForCheck();  
        }),
      )
      .subscribe(response => {
        console.warn(response);
        this.router.navigateByUrl(this.redirectTo);
      }, error => {
        console.warn(error);
        alert(this.translateService.translate('ERROR_INVALID_ID_OR_PASSWORD'));
      });
  }
}
