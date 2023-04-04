import {Component, createRef, render} from "preact";
import {WebTTY} from "./webtty";

interface MenuProps {
    tty: WebTTY
    // children: ComponentChildren;
    // buttons?: ComponentChildren;
    // title: string;
    // dismissHandler?: (hideModal?: () => void) => void;
}

export class Menu extends Component<MenuProps, {}> {
    ref = createRef<HTMLDivElement>();

    constructor() {
        super();
    }

    render() {
        return <div className="card" ref={this.ref}>
            <div className={"card-arrow"}>
                <span>{"<"}</span>
            </div>
            <div className="card-body">
                <h5 className="card-title">Menu</h5>
                <a href="/" target="_blank" className="card-link">New Session</a> <br/>
                <a href={`/session/create-readonly/${location.href.substring(location.href.lastIndexOf('/') + 1)}`} target="_blank" className="card-link">
                    Public Read-Only Session</a> <br/>
                <a href="#" className="card-link" onClick={() => {
                    this.props.tty.connection.close();
                }}>Close Session</a> <br/>
            </div>
        </div>;
    }
}

export function renderMenu(tty: WebTTY) {
    let newElem = document.createElement("div");
    document.body.prepend(newElem);
    newElem.setAttribute("class", "menu");

    render(<Menu tty={tty}> </Menu>, newElem);
}