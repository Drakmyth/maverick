import { GetModifyIWADModal, SaveModifiedIWAD } from "../../wailsjs/go/main/App";
import { Modal, ModalPosition } from "./modal";

export class ModifyIWADModal extends Modal<[string]> {
    override dialogId = "modify-iwad-dialog";
    override position = ModalPosition.CenterScreen;
    override coverClickDismiss = false;

    override async getModalContent(iwadId: string): Promise<string> {
        return await GetModifyIWADModal(iwadId);
    }

    public async submit(event: MouseEvent | SubmitEvent, iwadId: string): Promise<void> {
        event.preventDefault();

        let nameInput = document.getElementById("iwad-name-txt") as HTMLInputElement;
        let pathInput = document.getElementById("iwad-file-txt") as HTMLInputElement;

        await SaveModifiedIWAD(nameInput.value, pathInput.value, iwadId);

        this.close();
        // TODO: Add row to table instead of reloading page
        window.navigateTo("iwads");
    }
}
