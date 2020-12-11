interface Applicant {
  id: number;
  name: string;
  position: string;
  email: string;
  applyDate: number;
  status: 'examination' | 'progress' | 'end' | 'onboard';
  pass: boolean;
  progressDetail?: {
    label: string;
    complete: boolean;
  },
}

interface Paragraph {
  title: string;
  texts: string[];
}

interface JobPosition {
  id: number;
  category: string;
  name: string;
  summary: string;
  descriptions: Paragraph[];
}
