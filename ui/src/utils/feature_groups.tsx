//import MonitorIcon from '@mui/icons-material/Monitor';
import HomeIcon from '@mui/icons-material/Home';
import SmartToyIcon from '@mui/icons-material/SmartToy';
import { AiFillGithub, AiOutlineHome } from 'react-icons/ai';
import {FaXTwitter} from "react-icons/fa6";
import {BsLinkedin, BsGithub, BsDiscord, BsRobot} from "react-icons/bs";

export const products = [
    {
        "name": "Home",
        "url": "/",
        "icon": <AiOutlineHome className="text-sky-700" size={32} />

    },
    {
        "name": "Agents",
        "url": "/agents",
        "icon": <BsRobot className="text-sky-700"  size={32}/>
    },
]

export const socials = [
    {
        "name": "LinkedIn",
        "url": "https://www.linkedin.com/company/xpuls-ai",
        "icon": <BsLinkedin  size={32} color={"#0072b1"}/>
    },
    {
        "name": "Github",
        "url": "https://github.com/xpuls-labs",
        "icon": <BsGithub size={32} color={"black"}/>
    },
    {
        "name": "Twitter",
        "url": "https://x.com/xpulsai",
        "icon": <FaXTwitter size={32} color={"black"}/>
    },
    {
        "name": "Discord",
        "url": "https://discord.gg/AAS326HMQK",
        "icon": <BsDiscord size={32} color={"#7289DA"}/>
    },
]
