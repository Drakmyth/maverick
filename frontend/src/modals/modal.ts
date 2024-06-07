export enum ModalPosition {
    CenterScreen,
    Contextual,
}

export abstract class Modal<T extends unknown[]> {
    protected abstract dialogId: string;
    protected abstract position: ModalPosition;
    protected abstract coverClickDismiss: boolean;

    protected abstract getModalContent(...args: T): Promise<string>;

    public async open(event: MouseEvent, ...args: T) {
        this.closeAllModals();

        let template = document.createElement("template");
        template.innerHTML = await this.getModalContent(...args);
        let dialog = template.content.children[0] as HTMLDialogElement;

        if (this.coverClickDismiss) {
            dialog.onmousedown = this.getDialogCoverClickHandler(dialog, () => this.close());
        }

        let app = document.getElementById("app") as HTMLDivElement;
        app.append(dialog);

        if (this.position === ModalPosition.Contextual) {
            let margin = this.getContextualMargin(event.target as HTMLElement, dialog);
            dialog.style.marginLeft = `${margin.x}px`;
            dialog.style.marginTop = `${margin.y}px`;

            window.addEventListener("resize", function resized() {
                this.close();
                window.removeEventListener("resize", resized);
            });
        }

        dialog.showModal();
    }

    public submit(_event: MouseEvent | SubmitEvent, ..._args: T): void {}

    public close() {
        let dialog = document.getElementById(this.dialogId) as HTMLDialogElement;
        dialog.close();
        dialog.onmousedown = null;
        dialog.remove();
    }

    private closeAllModals() {
        let dialogs = document.getElementsByTagName("dialog");
        Array.from(dialogs).forEach((dialog) => {
            dialog.close();
            dialog.onmousedown = null;
            dialog.remove();
        });
    }

    private getDialogCoverClickHandler(dialog: HTMLDialogElement, callback: Function) {
        return (event: MouseEvent) => {
            let rect = dialog.getBoundingClientRect();
            let clickedInDialog =
                event.clientX >= rect.left &&
                event.clientX <= rect.right &&
                event.clientY >= rect.top &&
                event.clientY <= rect.bottom;

            if (!clickedInDialog) {
                callback();
            }
        };
    }

    private getContextualMargin(trigger: HTMLElement, dialog: HTMLDialogElement) {
        let triggerBounds = trigger.getBoundingClientRect();
        let dialogBounds = dialog.getBoundingClientRect();

        let viewportBounds = document.documentElement.getBoundingClientRect();

        let left = triggerBounds.right - dialogBounds.width;
        if (left < 0) {
            left = triggerBounds.left;
        }

        let top = triggerBounds.bottom;
        if (top + dialogBounds.height > viewportBounds.bottom) {
            top = triggerBounds.top - dialogBounds.height;
        }

        return new DOMPoint(left, top);
    }
}
