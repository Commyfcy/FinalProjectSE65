import * as React from 'react';
import AppBar from '@mui/material/AppBar';
 import Box from '@mui/material/Box';
 import Toolbar from '@mui/material/Toolbar';
 import Typography from '@mui/material/Typography';
 import Button from '@mui/material/Button';
 import IconButton from '@mui/material/IconButton';

 import MenuIcon from '@mui/icons-material/Menu';
 import AccountCircle from '@mui/icons-material/AccountCircle';
 import Switch from '@mui/material/Switch';
 import FormControlLabel from '@mui/material/FormControlLabel';
 import FormGroup from '@mui/material/FormGroup';
 import MenuItem from '@mui/material/MenuItem';
 import Menu from '@mui/material/Menu';
 import Avatar from '@mui/material/Avatar';
 import Container from '@mui/material/Container';
 import Tooltip from '@mui/material/Tooltip';
 import AdbIcon from '@mui/icons-material/Adb';

 import { useSession, signIn, signOut } from "next-auth/react"


export default function Component() {


  const pages = ['Products', 'Pricing', 'Blog'];

  const [anchorEl, setAnchorEl] = React.useState(null);

  const handleChange = (event) => {
    setAuth(event.target.checked);
  };

  const handleMenu = (event) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
  setAnchorEl(null);
  };

  const { data: session } = useSession()
  console.log(session);

  const [anchorElNav, setAnchorElNav] = React.useState(null);
  const [anchorElUser, setAnchorElUser] = React.useState(null);

  const handleOpenNavMenu = (event) => {
    setAnchorElNav(event.currentTarget);
  };
  const handleOpenUserMenu = (event) => {
    setAnchorElUser(event.currentTarget);
  };

  const handleCloseNavMenu = () => {
    setAnchorElNav(null);
  };

  const handleCloseUserMenu = () => {
  setAnchorElUser(null);
  };

  return (
    <div>
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
      <Toolbar disableGutters>
          <AdbIcon sx={{ display: { xs: 'none', md: 'flex' }, mr: 1 }} />
          <Typography
            variant="h6"
            noWrap
            component="a"
            href="/"
            sx={{
              mr: 2,
              display: { xs: 'none', md: 'flex' },
              fontFamily: 'monospace',
              fontWeight: 700,
              letterSpacing: '.3rem',
              color: 'inherit',
              textDecoration: 'none',
            }}
          >
            Logo
          </Typography>
          <Box sx={{ flexGrow: 1, display: { xs: 'none', md: 'flex' } }}>
            {pages.map((page) => (
              <Button
                key={page}
                onClick={handleCloseNavMenu}
                sx={{ my: 2, color: 'white', display: 'block' }}
              >
                {page}
              </Button>
            ))}
          </Box>
          {
          session
            ?<div>
          <IconButton
              size="large"
              aria-label="account of current user"
              aria-controls="menu-appbar"
              aria-haspopup="true"
              onClick={handleMenu}
              color="inherit"
           >
           <AccountCircle />
           </IconButton>
           <Menu
             id="menu-appbar"
             anchorEl={anchorEl}
             anchorOrigin={{
             vertical: 'top',
             horizontal: 'right',
          }}
          keepMounted
          transformOrigin={{
            vertical: 'top',
            horizontal: 'right',
          }}
          open={Boolean(anchorEl)}
          onClose={handleClose}
        >
          <MenuItem onClick={handleClose}> My Profile</MenuItem>
          <MenuItem onClick={()=> signOut()}>Logout</MenuItem>
        </Menu>
        </div>
        : <Button onClick={()=> signIn()}
color="inherit">Login</Button>}
        </Toolbar>
      </AppBar>
    </Box>
    </div>
)
}