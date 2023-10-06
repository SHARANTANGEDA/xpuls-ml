import Tooltip from '@mui/material/Tooltip';
import { styled } from '@mui/material/styles';


export const WhiteBackgroundTooltip = styled(Tooltip)(({ theme }) => ({
    tooltip: {
        backgroundColor: 'white',
        color: 'black',
        boxShadow: theme.shadows[1],
        fontSize: 12,
    },
}));