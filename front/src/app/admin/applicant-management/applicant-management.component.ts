import {
  ChangeDetectionStrategy,
  ChangeDetectorRef,
  Component,
  OnInit,
} from '@angular/core';
import { ApiService } from 'app/core';
import * as RxOp from 'rxjs/operators';
import { TranslateService } from 'app/core/translate.service';

enum ProcessTab {
  EXAMINATION,
  PROGRESS,
  END,
  ONBOARD,
}

@Component({
  selector: 'dg-applicant-management',
  templateUrl: './applicant-management.component.html',
  styleUrls: [ './applicant-management.component.scss' ],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ApplicantManagementComponent implements OnInit {
  public readonly ProcessTab: typeof ProcessTab = ProcessTab;
  public readonly tabs = [
    { category: ProcessTab.EXAMINATION, label: this.translateService.translate('EXAMINATING_APPLICATION') },
    { category: ProcessTab.PROGRESS, label: this.translateService.translate('SCREENING') },
    { category: ProcessTab.END, label: this.translateService.translate('END_SCREEN') },
    { category: ProcessTab.ONBOARD, label: this.translateService.translate('ONBOARDING') },
  ];
  public activeTab: ProcessTab = ProcessTab.EXAMINATION;
  public isFetching: boolean = false;

  public selectedMoreWaitedApplicants: Applicant[] = [];
  public selectedOtherApplicants: Applicant[] = [];
  public selectedProgressApplicants: Applicant[] = [];
  public selectedIncompletedProgressApplicants: Applicant[] = [];
  public selectedOnboardingApplicants: Applicant[] = [];

  private readonly currentDate: number = Number(new Date());
  private applicants: Applicant[] = [];

  public get moreWaitedExaminationApplicants(): Applicant[] {
    return this.examinationApplicants.filter(applicant => (this.currentDate - applicant.applyDate) > 60000 * 60 * 24 * 5);
  }

  public get otherExaminationApplicants(): Applicant[] {
    return this.examinationApplicants.filter(applicant => (this.currentDate - applicant.applyDate) <= 60000 * 60 * 24 * 5);
  }

  public get progressCompletedApplicants(): Applicant[] {
    return this.progressApplicants.filter(applicant => applicant.progressDetail && applicant.progressDetail.complete);
  }

  public get onProgressApplicants(): Applicant[] {
    return this.progressApplicants.filter(applicant => applicant.progressDetail && !applicant.progressDetail.complete);
  }

  public get passedApplicants(): Applicant[] {
    return this.endedApplicants.filter(applicant => applicant.pass);
  }

  public get failedApplicants(): Applicant[] {
    return this.endedApplicants.filter(applicant => !applicant.pass);
  }

  public get onboardingApplicants(): Applicant[] {
    return this.applicants.filter(applicant => applicant.status === 'onboard');
  }

  private get examinationApplicants(): Applicant[] {
    return this.applicants.filter(applicant => applicant.status === 'examination');
  }

  private get progressApplicants(): Applicant[] {
    return this.applicants.filter(applicant => applicant.status === 'progress');
  }

  private get endedApplicants(): Applicant[] {
    return this.applicants.filter(applicant => applicant.status === 'end');
  }

  public constructor(
    private readonly apiService: ApiService,
    private readonly changeDetectorRef: ChangeDetectorRef,
    private readonly translateService: TranslateService,
  ) {
    return;
  }

  public ngOnInit(): void {
    this.fetchApplicants();
  }

  public getNumberOfApplicant(tab: ProcessTab): number {
    switch (tab) {
      case ProcessTab.EXAMINATION:
        return this.applicants.filter(applicant => applicant.status === 'examination').length;
      case ProcessTab.PROGRESS:
        return this.applicants.filter(applicant => applicant.status === 'progress').length;
      case ProcessTab.END:
        return this.applicants.filter(applicant => applicant.status === 'end').length;
      case ProcessTab.ONBOARD:
        return this.applicants.filter(applicant => applicant.status === 'onboard').length;
      default:
        return 0;
    }
  }

  public beautifyDate(timestamp: number): string {
    const date = new Date(timestamp);
    return `${date.getFullYear()}.${(date.getMonth() + 1).toString(10).padStart(2, '0')}.${date.getDate().toString(10).padStart(2, '0')}`;
  }

  public isSelectedMoreWaitedExaminationApplicant(applicant: Applicant): boolean {
    return this.selectedMoreWaitedApplicants.indexOf(applicant) > -1;
  }

  public onClickMoreWaitedExaminationApplicant(applicant: Applicant): void {
    const index = this.selectedMoreWaitedApplicants.indexOf(applicant);
    if (index > -1) {
      this.selectedMoreWaitedApplicants.splice(index, 1);
    } else {
      this.selectedMoreWaitedApplicants.push(applicant);
    }
  }

  public isSelectedOtherExaminationApplicant(applicant: Applicant): boolean {
    return this.selectedOtherApplicants.indexOf(applicant) > -1;
  }

  public onClickOtherExaminationApplicant(applicant: Applicant): void {
    const index = this.selectedOtherApplicants.indexOf(applicant);
    if (index > -1) {
      this.selectedOtherApplicants.splice(index, 1);
    } else {
      this.selectedOtherApplicants.push(applicant);
    }
  }

  public isSelectedProgressApplicant(applicant: Applicant): boolean {
    return this.selectedProgressApplicants.indexOf(applicant) > -1;
  }

  public onClickProgressApplicant(applicant: Applicant): void {
    const index = this.selectedProgressApplicants.indexOf(applicant);
    if (index > -1) {
      this.selectedProgressApplicants.splice(index, 1);
    } else {
      this.selectedProgressApplicants.push(applicant);
    }
  }

  public isSelectedIncompletedProgressApplicant(applicant: Applicant): boolean {
    return this.selectedIncompletedProgressApplicants.indexOf(applicant) > -1;
  }

  public onClickIncompletedProgressApplicant(applicant: Applicant): void {
    const index = this.selectedIncompletedProgressApplicants.indexOf(applicant);
    if (index > -1) {
      this.selectedIncompletedProgressApplicants.splice(index, 1);
    } else {
      this.selectedIncompletedProgressApplicants.push(applicant);
    }
  }

  public isSelectedOnboardingApplicant(applicant: Applicant): boolean {
    return this.selectedOnboardingApplicants.indexOf(applicant) > -1;
  }

  public onClickOnboardApplicant(applicant: Applicant): void {
    const index = this.selectedOnboardingApplicants.indexOf(applicant);
    if (index > -1) {
      this.selectedOnboardingApplicants.splice(index, 1);
    } else {
      this.selectedOnboardingApplicants.push(applicant);
    }
  }

  public onMoreWaitedExaminationApplicantPassButton(): void {
    if (this.selectedMoreWaitedApplicants.length === 0) {
      alert(this.translateService.translate('ERROR_NO_SELECTED_APPLICANT'));
      return;
    }
    if (confirm(this.translateService.translate('CONFIRM_PASS_APPLICANT_MESSAGE').replace('$1', `${this.selectedMoreWaitedApplicants.length}`))) {
      const screeningLabel = prompt(this.translateService.translate('ENTER_SCREENING_STEP_NAME'));
      if (!screeningLabel) {
        if (typeof screeningLabel === 'string') {
          alert(this.translateService.translate('ERROR_SCREENING_STEP_NAME_CANNOT_BE_BLANK'));
        }
        return;
      }
      this.selectedMoreWaitedApplicants.forEach(applicant => {
        applicant.status = 'progress';
        applicant.progressDetail = {
          label: screeningLabel,
          complete: false,
        };
      });
      this.updateApplicants(this.selectedMoreWaitedApplicants);
    }
  }

  public onMoreWaitedExaminationApplicantFailButton(): void {
    if (this.selectedMoreWaitedApplicants.length === 0) {
      alert(this.translateService.translate('ERROR_NO_SELECTED_APPLICANT'));
      return;
    }
    if (confirm(this.translateService.translate('CONFIRM_FAILURE_APPLICANT_MESSAGE').replace('$1', `${this.selectedMoreWaitedApplicants.length}`))) {
      this.selectedMoreWaitedApplicants.forEach(applicant => applicant.status = 'end');
      this.updateApplicants(this.selectedMoreWaitedApplicants);
    }
  }

  public onOtherExaminationApplicantPassButton(): void {
    if (this.selectedOtherApplicants.length === 0) {
      alert(this.translateService.translate('ERROR_NO_SELECTED_APPLICANT'));
      return;
    }
    if (confirm(this.translateService.translate('CONFIRM_PASS_APPLICANT_MESSAGE').replace('$1', `${this.selectedOtherApplicants.length}`))) {
      const screeningLabel = prompt(this.translateService.translate('ENTER_SCREENING_STEP_NAME'));
      if (!screeningLabel) {
        if (typeof screeningLabel === 'string') {
          alert(this.translateService.translate('ERROR_SCREENING_STEP_NAME_CANNOT_BE_BLANK'));
        }
        return;
      }
      this.selectedOtherApplicants.forEach(applicant => {
        applicant.status = 'progress';
        applicant.progressDetail = {
          label: screeningLabel,
          complete: false,
        };
      });
      this.updateApplicants(this.selectedOtherApplicants);
    }
  }

  public onOtherExaminationApplicantFailButton(): void {
    if (this.selectedOtherApplicants.length === 0) {
      alert(this.translateService.translate('ERROR_NO_SELECTED_APPLICANT'));
      return;
    }
    if (confirm(this.translateService.translate('CONFIRM_FAILURE_APPLICANT_MESSAGE').replace('$1', `${this.selectedOtherApplicants.length}`))) {
      this.selectedOtherApplicants.forEach(applicant => applicant.status = 'end');
      this.updateApplicants(this.selectedOtherApplicants);
    }
  }

  public onProgressApplicantPassButton(): void {
    if (this.selectedProgressApplicants.length === 0) {
      alert(this.translateService.translate('ERROR_NO_SELECTED_APPLICANT'));
      return;
    }
    if (confirm(this.translateService.translate('CONFIRM_PASS_APPLICANT_MESSAGE').replace('$1', `${this.selectedProgressApplicants.length}`))) {
      if (confirm(this.translateService.translate('CONTINUE_SCREENING'))) {
        this.selectedProgressApplicants.forEach(applicant => applicant.status = 'onboard');
      } else {
        const screeningLabel = prompt(this.translateService.translate('ENTER_SCREENING_STEP_NAME'));
        if (!screeningLabel) {
          if (typeof screeningLabel === 'string') {
            alert(this.translateService.translate('ERROR_SCREENING_STEP_NAME_CANNOT_BE_BLANK'));
          }
          return;
        }
        this.selectedProgressApplicants.forEach(applicant => applicant.progressDetail.label = screeningLabel);
      }
      this.updateApplicants(this.selectedProgressApplicants);
    }
  }

  public onProgressApplicantFailButton(): void {
    if (this.selectedProgressApplicants.length === 0) {
      alert(this.translateService.translate('ERROR_NO_SELECTED_APPLICANT'));
      return;
    }
    if (confirm(this.translateService.translate('CONFIRM_FAILURE_APPLICANT_MESSAGE').replace('$1', `${this.selectedProgressApplicants.length}`))) {
      this.selectedProgressApplicants.forEach(applicant => applicant.status = 'end');
      this.updateApplicants(this.selectedProgressApplicants);
    }
  }

  public onClickIncompletedProgressApplicantPassButton(): void {
    if (this.selectedIncompletedProgressApplicants.length === 0) {
      alert(this.translateService.translate('ERROR_NO_SELECTED_APPLICANT'));
      return;
    }
    if (confirm(this.translateService.translate('CONFIRM_PASS_APPLICANT_MESSAGE').replace('$1', `${this.selectedIncompletedProgressApplicants.length}`))) {
      if (confirm(this.translateService.translate('CONTINUE_SCREENING'))) {
        this.selectedIncompletedProgressApplicants.forEach(applicant => applicant.status = 'onboard');
      } else {
        const screeningLabel = prompt(this.translateService.translate('ENTER_SCREENING_STEP_NAME'));
        if (!screeningLabel) {
          if (typeof screeningLabel === 'string') {
            alert(this.translateService.translate('ERROR_SCREENING_STEP_NAME_CANNOT_BE_BLANK'));
          }
          return;
        }
        this.selectedIncompletedProgressApplicants.forEach(applicant => applicant.progressDetail.label = screeningLabel);
      }
      this.updateApplicants(this.selectedIncompletedProgressApplicants);
    }
  }

  public onClickIncompletedProgressApplicantFailButton(): void {
    if (this.selectedIncompletedProgressApplicants.length === 0) {
      alert(this.translateService.translate('ERROR_NO_SELECTED_APPLICANT'));
      return;
    }
    if (confirm(this.translateService.translate('CONFIRM_FAILURE_APPLICANT_MESSAGE').replace('$1', `${this.selectedIncompletedProgressApplicants.length}`))) {
      this.selectedIncompletedProgressApplicants.forEach(applicant => applicant.status = 'end');
      this.updateApplicants(this.selectedIncompletedProgressApplicants);
    }
  }

  public onboardingApplicantPassButton(): void {
    if (this.selectedOnboardingApplicants.length === 0) {
      alert(this.translateService.translate('ERROR_NO_SELECTED_APPLICANT'));
      return;
    }
    if (confirm(this.translateService.translate('CONFIRM_PASS_APPLICANT_MESSAGE').replace('$1', `${this.selectedOnboardingApplicants.length}`))) {
      this.selectedOnboardingApplicants.forEach(applicant => {
        applicant.status = 'end';
        applicant.pass = true;
      });
      this.updateApplicants(this.selectedOnboardingApplicants);
    }
  }

  public onboardingApplicantFailButton(): void {
    if (this.selectedOnboardingApplicants.length === 0) {
      alert(this.translateService.translate('ERROR_NO_SELECTED_APPLICANT'));
    }
    if (confirm(this.translateService.translate('CONFIRM_FAILURE_APPLICANT_MESSAGE').replace('$1', `${this.selectedOnboardingApplicants.length}`))) {
      this.selectedOnboardingApplicants.forEach(applicant => applicant.status = 'end');
      this.updateApplicants(this.selectedOnboardingApplicants);
    }
  }

  private fetchApplicants(): void {
    this.isFetching = true;
    this.apiService.getApplicants()
      .pipe(
        RxOp.finalize(() => this.isFetching = false),
      )
      .subscribe(applicants => {
        this.selectedMoreWaitedApplicants.splice(0);
        this.selectedOtherApplicants.splice(0);
        this.selectedProgressApplicants.splice(0);
        this.selectedIncompletedProgressApplicants.splice(0);
        this.selectedOnboardingApplicants.splice(0);
        this.applicants = applicants;
        this.changeDetectorRef.markForCheck();
      }, err => {
        alert(this.translateService.translate('ERROR_FETCH_APPLICANT_FAILED'));
      });
  }

  private updateApplicants(applicants: Applicant[]): void {
    this.apiService.updateApplicants(applicants)
      .subscribe(() => {
        this.fetchApplicants();
      }, err => {
        alert(this.translateService.translate('ERROR_SAVE_APPLICANT_CHANGE_FAIELD'));
      });
  }
}
