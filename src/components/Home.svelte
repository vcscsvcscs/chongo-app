<script>
    import { onMount } from "svelte";
    import { Link } from 'svelte-routing';
    import {mainsocket} from '../socketcom'
    import {UserLogedin} from '../Authenticate'
    import Dashboard from "../components/Home/Dashboard.svelte";
    import Chat from "../components/Home/Chat.svelte";
    import Discover from "../components/Home/Discover.svelte";
    import Pagetopbutton from "../components/Pagetopbutton.svelte"; 
    let switchmenu = "dashboard";
    onMount(()=>{
        $mainsocket = {"token":(UserLogedin()?UserLogedin():"close"),"contents":[],"types":[]};
    });
    mainsocket.subscribe(
        (value) => {
            console.log(value);
        }
    );
</script>

<link rel="stylesheet" href="assets/css/styles.css" />
<div id="wrapper">
    <nav
        class="navbar navbar-dark align-items-start sidebar sidebar-dark accordion bg-gradient-primary p-0" style=" min-width: 20vw;">
        <div class="container-fluid d-flex flex-column p-0">
            <Link
                class="navbar-brand d-flex justify-content-center align-items-center sidebar-brand m-0"
                to="/home"
                ><img
                    src="assets/img/logo.png"
                    style="width: 50px;height: 50px;"
                    alt="Logo"
                />
                <div class="sidebar-brand-icon rotate-n-15" />
                <div class="sidebar-brand-text mx-3"><span>Chongo</span></div>
            </Link>
            <hr class="sidebar-divider my-0" />
            <ul class="navbar-nav text-light" id="accordionSidebar">
                <li class="nav-item">
                    <button class="nav-link active" on:click="{()=>{switchmenu = 'dashboard'}}"
                        ><i class="fas fa-tachometer-alt" /><span
                            >Dashboard</span
                        ></button
                    >
                </li>
                <li class="nav-item"><button on:click="{()=>{switchmenu = 'discover'}}" type="button" class="nav-link"><i class="fas fa-table"></i><span>Discover</span></button></li>
                <li class="nav-item"><button type="button" class="nav-link" on:click="{()=>{switchmenu = 'chat'}}"><i class="fa fa-wechat" /><span>Chat</span></button></li>
            </ul>
            <div class="text-center d-none d-md-inline">
                <button
                    class="btn rounded-circle border-0"
                    id="sidebarToggle"
                    type="button"
                />
            </div>
        </div>
    </nav>
    <div class="d-flex flex-column" id="content-wrapper" style=" min-width: 80vw;">
        <div id="content" style="height: 96vh;">
            <nav
                class="navbar navbar-light navbar-expand bg-white shadow mb-4 topbar static-top">
                <div class="container-fluid">
                    <button
                        class="btn btn-link d-md-none rounded-circle me-3"
                        id="sidebarToggleTop"
                        type="button"><i class="fas fa-bars" /></button
                    >
                    <form
                        class="d-none d-sm-inline-block me-auto ms-md-3 my-2 my-md-0 mw-100 navbar-search"
                    >
                        <div class="input-group">
                            <input
                                class="bg-light form-control border-0 small"
                                type="text"
                                placeholder="Search for ..."
                            /><button class="btn btn-primary py-0" type="button"
                                ><i class="fas fa-search" /></button
                            >
                        </div>
                    </form>
                    <ul class="navbar-nav flex-nowrap ms-auto">
                        <li class="nav-item dropdown d-sm-none no-arrow">
                            <a
                                class="dropdown-toggle nav-link"
                                aria-expanded="false"
                                data-bs-toggle="dropdown"
                                href="/"><i class="fas fa-search" /></a
                            >
                            <div
                                class="dropdown-menu dropdown-menu-end p-3 animated--grow-in"
                                aria-labelledby="searchDropdown"
                            >
                                <form class="me-auto navbar-search w-100">
                                    <div class="input-group">
                                        <input
                                            class="bg-light form-control border-0 small"
                                            type="text"
                                            placeholder="Search for ..."
                                        />
                                        <div class="input-group-append">
                                            <button
                                                class="btn btn-primary py-0"
                                                type="button"
                                                ><i
                                                    class="fas fa-search"
                                                /></button
                                            >
                                        </div>
                                    </div>
                                </form>
                            </div>
                        </li>
                        <li class="nav-item dropdown no-arrow mx-1">
                            <div class="nav-item dropdown no-arrow">
                                <a
                                    class="dropdown-toggle nav-link"
                                    aria-expanded="false"
                                    data-bs-toggle="dropdown"
                                    href="/"
                                    ><span class="badge bg-danger badge-counter"
                                        >3+</span
                                    ><i class="fas fa-bell fa-fw" /></a
                                >
                                <div
                                    class="dropdown-menu dropdown-menu-end dropdown-list animated--grow-in"
                                >
                                    <h6 class="dropdown-header">
                                        alerts center
                                    </h6>
                                    <a
                                        class="dropdown-item d-flex align-items-center"
                                        href="/"
                                    >
                                        <div class="me-3">
                                            <div class="bg-primary icon-circle">
                                                <i
                                                    class="fas fa-file-alt text-white"
                                                />
                                            </div>
                                        </div>
                                        <div>
                                            <span class="small text-gray-500"
                                                >December 12, 2019</span
                                            >
                                            <p>
                                                A new monthly report is ready to
                                                download!
                                            </p>
                                        </div>
                                    </a><a
                                        class="dropdown-item d-flex align-items-center"
                                        href="/"
                                    >
                                        <div class="me-3">
                                            <div class="bg-success icon-circle">
                                                <i
                                                    class="fas fa-donate text-white"
                                                />
                                            </div>
                                        </div>
                                        <div>
                                            <span class="small text-gray-500"
                                                >December 7, 2019</span
                                            >
                                            <p>
                                                $290.29 has been deposited into
                                                your account!
                                            </p>
                                        </div>
                                    </a><a
                                        class="dropdown-item d-flex align-items-center"
                                        href="/"
                                    >
                                        <div class="me-3">
                                            <div class="bg-warning icon-circle">
                                                <i
                                                    class="fas fa-exclamation-triangle text-white"
                                                />
                                            </div>
                                        </div>
                                        <div>
                                            <span class="small text-gray-500"
                                                >December 2, 2019</span
                                            >
                                            <p>
                                                Spending Alert: We've noticed
                                                unusually high spending for your
                                                account.
                                            </p>
                                        </div>
                                    </a><a
                                        class="dropdown-item text-center small text-gray-500"
                                        href="/">Show All Alerts</a
                                    >
                                </div>
                            </div>
                        </li>
                        <li class="nav-item dropdown no-arrow mx-1">
                            <div class="nav-item dropdown no-arrow">
                                <a
                                    class="dropdown-toggle nav-link"
                                    aria-expanded="false"
                                    data-bs-toggle="dropdown"
                                    href="/"
                                    ><span class="badge bg-danger badge-counter"
                                        >7</span
                                    ><i class="fas fa-envelope fa-fw" /></a
                                >
                                <div
                                    class="dropdown-menu dropdown-menu-end dropdown-list animated--grow-in"
                                >
                                    <h6 class="dropdown-header">
                                        alerts center
                                    </h6>
                                    <a
                                        class="dropdown-item d-flex align-items-center"
                                        href="/"
                                    >
                                        <div class="dropdown-list-image me-3">
                                            <img
                                                class="rounded-circle"
                                                src="assets/img/avatars/avatar4.jpeg"
                                                alt="Avatar 4"
                                            />
                                            <div
                                                class="bg-success status-indicator"
                                            />
                                        </div>
                                        <div class="fw-bold">
                                            <div class="text-truncate">
                                                <span
                                                    >Hi there! I am wondering if
                                                    you can help me with a
                                                    problem I've been having.</span
                                                >
                                            </div>
                                            <p class="small text-gray-500 mb-0">
                                                Emily Fowler - 58m
                                            </p>
                                        </div>
                                    </a><a
                                        class="dropdown-item d-flex align-items-center"
                                        href="/"
                                    >
                                        <div class="dropdown-list-image me-3">
                                            <img
                                                class="rounded-circle"
                                                src="assets/img/avatars/avatar2.jpeg"
                                                alt="Avatar 2"
                                            />
                                            <div class="status-indicator" />
                                        </div>
                                        <div class="fw-bold">
                                            <div class="text-truncate">
                                                <span
                                                    >I have the photos that you
                                                    ordered last month!</span
                                                >
                                            </div>
                                            <p class="small text-gray-500 mb-0">
                                                Jae Chun - 1d
                                            </p>
                                        </div>
                                    </a><a
                                        class="dropdown-item d-flex align-items-center"
                                        href="/"
                                    >
                                        <div class="dropdown-list-image me-3">
                                            <img
                                                class="rounded-circle"
                                                src="assets/img/avatars/avatar3.jpeg"
                                                alt="Avatar 3"
                                            />
                                            <div
                                                class="bg-warning status-indicator"
                                            />
                                        </div>
                                        <div class="fw-bold">
                                            <div class="text-truncate">
                                                <span
                                                    >Last month's report looks
                                                    great, I am very happy with
                                                    the progress so far, keep up
                                                    the good work!</span
                                                >
                                            </div>
                                            <p class="small text-gray-500 mb-0">
                                                Morgan Alvarez - 2d
                                            </p>
                                        </div>
                                    </a><a
                                        class="dropdown-item d-flex align-items-center"
                                        href="/"
                                    >
                                        <div class="dropdown-list-image me-3">
                                            <img
                                                class="rounded-circle"
                                                src="assets/img/avatars/avatar5.jpeg"
                                                alt="Avatar 5"
                                            />
                                            <div
                                                class="bg-success status-indicator"
                                            />
                                        </div>
                                        <div class="fw-bold">
                                            <div class="text-truncate">
                                                <span
                                                    >Am I a good boy? The reason
                                                    I ask is because someone
                                                    told me that people say this
                                                    to all dogs, even if they
                                                    aren't good...</span
                                                >
                                            </div>
                                            <p class="small text-gray-500 mb-0">
                                                Chicken the Dog · 2w
                                            </p>
                                        </div>
                                    </a><a
                                        class="dropdown-item text-center small text-gray-500"
                                        href="/">Show All Alerts</a
                                    >
                                </div>
                            </div>
                            <div
                                class="shadow dropdown-list dropdown-menu dropdown-menu-end"
                                aria-labelledby="alertsDropdown"
                            />
                        </li>
                        <div class="d-none d-sm-block topbar-divider" />
                        <li class="nav-item dropdown no-arrow">
                            <div class="nav-item dropdown no-arrow">
                                <a
                                    class="dropdown-toggle nav-link"
                                    aria-expanded="false"
                                    data-bs-toggle="dropdown"
                                    href="/"
                                    ><span
                                        class="d-none d-lg-inline me-2 text-gray-600 small"
                                        >Valerie Luna</span
                                    ><img
                                        class="border rounded-circle img-profile"
                                        src="assets/img/avatars/avatar1.jpeg"
                                        alt="Avatar 1"
                                    /></a
                                >
                                <div
                                    class="dropdown-menu shadow dropdown-menu-end animated--grow-in"
                                >
                                    <a class="dropdown-item" href="/"
                                        ><i
                                            class="fas fa-user fa-sm fa-fw me-2 text-gray-400"
                                        />&nbsp;Profile</a
                                    ><a class="dropdown-item" href="/"
                                        ><i
                                            class="fas fa-cogs fa-sm fa-fw me-2 text-gray-400"
                                        />&nbsp;Settings</a
                                    ><a class="dropdown-item" href="/"
                                        ><i
                                            class="fas fa-list fa-sm fa-fw me-2 text-gray-400"
                                        />&nbsp;Activity log</a
                                    >
                                    <div class="dropdown-divider" />
                                    <a class="dropdown-item" href="/"
                                        ><i
                                            class="fas fa-sign-out-alt fa-sm fa-fw me-2 text-gray-400"
                                        />&nbsp;Logout</a
                                    >
                                </div>
                            </div>
                        </li>
                    </ul>
                </div>
            </nav>
            {#if switchmenu == "dashboard"}
                <Dashboard/>
            {:else if switchmenu == "chat"}
                <Chat/>
            {:else if switchmenu == "discover"}
                <Discover/>
            {/if}
        </div>
        <footer class="bg-white sticky-footer" style="height: 4vh;">
            <div class="container my-auto">
                <div class="text-center my-auto copyright" >
                    <span>Copyright © Chongo 2022</span>
                </div>
            </div>
        </footer>
    </div>
    <Pagetopbutton/>
</div>
