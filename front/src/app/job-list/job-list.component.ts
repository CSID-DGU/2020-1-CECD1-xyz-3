import {
  ChangeDetectionStrategy,
  ChangeDetectorRef,
  Component,
  Inject,
  OnInit,
} from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ApiService } from 'app/core';
import { TranslateService } from 'app/core/translate.service';

interface UserInfo {
  name: string;
  email: string;
}

@Component({
  selector: 'dg-job-list',
  templateUrl: './job-list.component.html',
  styleUrls: [ './job-list.component.scss' ],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class JobListComponent implements OnInit {
  public descriptiionTitle: string = this.translateService.translate('COMPANY_SAMPLE_TITLE');
  public descriptiionDetail: string = this.translateService.translate('COMPANY_SAMPLE_DESCRIPTION');
  public positions: JobPosition[] = [];

  public navigationMenus: any[] = [
    {
      label: this.translateService.translate('RECRUTING_POSITIONS'),
      link: '/position-list',
    },
    {
      label: this.translateService.translate('VIEW_MY_APPLICATION'),
      disabled: true,
    },
  ];
  public companyName: string;
  public fullCompanyName: string;
  public positionId: number;
  public applicantInfoModalVisible: boolean = false;

  public inputUserName: string;
  public inputEmailName: string;


  public constructor(
    @Inject('userName') public userInfo: UserInfo,
    private readonly apiService: ApiService,
    private readonly changeDetectorRef: ChangeDetectorRef,
    private readonly translateService: TranslateService,
    private readonly route: ActivatedRoute,
  ) {
    if (userInfo.name === undefined || userInfo.email === undefined) {
      this.applicantInfoModalVisible = true;
    }
    this.fullCompanyName = route.snapshot.paramMap.get('companyName');
    if (/^[a-zA-Z0-9]+\.[a-zA-Z0-9]+$/.test(this.fullCompanyName)) {
      this.companyName = this.fullCompanyName.substr(0, this.fullCompanyName.indexOf('.'));
      this.descriptiionDetail = this.descriptiionDetail.replace(/\$1/g, this.companyName);
    } else {
      alert(this.translateService.translate('ERROR_INVALID_PATH'));
      return;
    }
    const positionId = route.snapshot.paramMap.get('positionId');
    if (positionId) {
      this.positionId = Number(positionId);
    }
  }

  public ngOnInit(): void {
    this.fetchPositions();
  }

  public getPosition(id: number): JobPosition {
    return this.positions.find(position => position.id === id);
  }

  public onClickConfirmButton(): void {
    this.applicantInfoModalVisible = false;
    this.userInfo.name = this.inputUserName;
    this.userInfo.email = this.inputEmailName;
  }

  public onClickCancelButton(): void {
    this.applicantInfoModalVisible = false;
    this.userInfo.name = null;
    this.userInfo.email = null;
  }

  public onClickLoginButton(): void {
    this.applicantInfoModalVisible = true;
  }

  public onClickApplyButton(): void {
    const position = this.getPosition(this.positionId);
    if (position && confirm(this.translateService.translate('APPLY_CONFIRM_MESSAGE').replace('$1', position.name))) {
      this.apiService.appendApplicants({
        id: -1,
        applyDate: Number(new Date()),
        email: this.userInfo.email,
        name: this.userInfo.name,
        pass: false,
        position: position.name,
        status: 'examination',
      }, this.fullCompanyName)
        .subscribe(() => {
          alert(this.translateService.translate('APPLICATION_ACCEPTED'));
        }, error => {
          alert(this.translateService.translate('ERROR_APPLICATION_ACCEPTANCE_FAILED'));
        });
    }
  }

  private fetchPositions(): void {
    this.apiService.getPositions(this.fullCompanyName)
      .subscribe(positions => {
        this.positions = positions;
        this.changeDetectorRef.markForCheck();
      }, err => {
        alert(this.translateService.translate('ERROR_FETCH_POSITION_FAILED'));
      })
  }
}
