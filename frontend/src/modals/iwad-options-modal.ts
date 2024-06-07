import { GetIWADOptionsModal } from "../../wailsjs/go/main/App";
import { Modal, ModalPosition } from "./modal";

export class IWADOptionsModal extends Modal<[string]> {
    override dialogId = "iwad-options-modal";
    override position = ModalPosition.Contextual;
    override coverClickDismiss = true;

    override async getModalContent(iwadId: string): Promise<string> {
        return await GetIWADOptionsModal(iwadId);
    }
}
