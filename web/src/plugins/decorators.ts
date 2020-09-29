export function Confirm(
  caption: string,
  title = 'Confirm',
  cancelFunction?: Function
) {
  return function(
    target: any,
    propertyKey: string,
    descriptor: PropertyDescriptor
  ) {
    const orgValue: Function = descriptor.value;

    descriptor.value = function(...args: any) {
      const v = target.$q
        .dialog({
          title: title,
          message: caption,
          persistent: true,
          cancel: 'Cancel',
          ok: 'OK',
        })
        .onOk(() => orgValue.apply(this, args));
      if (cancelFunction) {
        v.onCancel(cancelFunction);
      }
    };

    return descriptor;
  };
}
