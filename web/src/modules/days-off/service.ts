import { CrudService } from '@/services/crud';
import { DayOff } from '@/types/logic';

export class DayOffService extends CrudService<DayOff> {
  constructor() {
    super('/days-off');
  }
}
