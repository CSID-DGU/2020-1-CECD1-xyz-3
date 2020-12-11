import { Injectable } from '@angular/core';
import {
  ActivatedRouteSnapshot,
  CanActivate,
  Router,
  RouterStateSnapshot,
} from '@angular/router';
import { ApiService } from 'app/core';
import * as Rx from 'rxjs';
import * as RxOp from 'rxjs/operators';

@Injectable()
export class LoginGuard implements CanActivate {
  public constructor(
    public apiService: ApiService,
    public router: Router,
  ) {
    return;
  }

  public canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): any {
    return this.apiService.isAuthenticated()
      .pipe(
        RxOp.map(() => true),
        RxOp.catchError(() => {
          this.router.navigate([ `/login` ], { queryParams: { redirectTo: state.url } });
          return Rx.of(false);
        }),
      );
  }
}
