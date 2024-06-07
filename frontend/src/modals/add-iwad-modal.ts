import { GetContent, SaveIWAD } from "../../wailsjs/go/main/App";
import { Modal, ModalPosition } from "./modal";

export class AddIWADModal extends Modal<[void]> {
    override dialogId = "add-iwad-dialog";
    override position = ModalPosition.CenterScreen;
    override coverClickDismiss = false;

    override async getModalContent(): Promise<string> {
        return await GetContent("add-iwad-modal");
    }

    override async submit(event: MouseEvent | SubmitEvent): Promise<void> {
        event.preventDefault();

        let nameInput = document.getElementById("iwad-name-txt") as HTMLInputElement;
        let pathInput = document.getElementById("iwad-file-txt") as HTMLInputElement;

        await SaveIWAD(nameInput.value, pathInput.value);

        this.close();
        // TODO: Add row to table instead of reloading page
        window.navigateTo("iwads");
    }
}
