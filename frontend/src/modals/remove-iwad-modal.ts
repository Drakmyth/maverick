import { GetRemoveIWADModal, RemoveIWAD } from "../../wailsjs/go/main/App";
import { Modal, ModalPosition } from "./modal";

export class RemoveIWADModal extends Modal<[string]> {
    override dialogId = "remove-iwad-dialog";
    override position = ModalPosition.CenterScreen;
    override coverClickDismiss = false;

    override async getModalContent(iwadId: string): Promise<string> {
        return await GetRemoveIWADModal(iwadId);
    }

    public async submit(_event: MouseEvent | SubmitEvent, iwadId: string): Promise<void> {
        let success = await RemoveIWAD(iwadId);
        if (!success) {
            return;
        }

        this.close();
        // TODO: Remove row from table instead of reloading page
        window.navigateTo("iwads");
    }
}
