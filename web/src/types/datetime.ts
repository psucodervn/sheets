import formatter from '@/utils/formatter';

interface Month {
  month: number;
  year: number;
}

class TimeRange {
  from: Date;
  to: Date;

  constructor(from: Date, to: Date) {
    this.from = from;
    this.to = to;
  }

  get label(): string {
    return `${formatter.dateDDMMYYYY(this.from)} to ${formatter.dateDDMMYYYY(
      this.to,
    )}`;
  }
}

export { Month, TimeRange };
