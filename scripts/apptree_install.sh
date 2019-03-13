#!/bin/bash
{
    set -e
    SUDO=''
    if [ "$(id -u)" != "0" ]; then
      SUDO='sudo'
      echo "This script requires superuser access."
      echo "You will be prompted for your password by sudo."
      # clear any previous sudo permission
      sudo -k
    fi


    # run inside sudo
    $SUDO bash <<SCRIPT
  set -e

  echoerr() { echo "\$@" 1>&2; }

  if [[ ! ":\$PATH:" == *":/usr/local/bin:"* ]]; then
    echoerr "Your path is missing /usr/local/bin, you need to add this to use this installer."
    exit 1
  fi

  if [ "\$(uname)" == "Darwin" ]; then
    OS=darwin
    URL="https://storage.googleapis.com/apptreeworkflow/binaries/apptree_darwin_amd64"
    echo "OS is darwin. Download URL is \$URL"
  elif [ "\$(expr substr \$(uname -s) 1 5)" == "Linux" ]; then
    OS=linux
    URL="https://storage.googleapis.com/apptreeworkflow/binaries/apptree_linux_amd64"
    echo "OS is linux. Download URL is \$URL"
  else
    echoerr "This installer is only supported on Linux and MacOS"
    exit 1
  fi

  ARCH="\$(uname -m)"
  if [ "\$ARCH" == "x86_64" ]; then
    ARCH=x64
  elif [[ "\$ARCH" == arm* ]]; then
    ARCH=arm
  else
    echoerr "unsupported arch: \$ARCH"
    exit 1
  fi

  #mkdir -p /usr/local/lib
  echo AppTree cache, cert, key files will be installed under HOME/.apptree
  echo Creating ".apptree" directory
  mkdir -p $HOME/.apptree
  chmod 666 $HOME/.apptree
  cd /usr/local/bin
  rm -rf apptree
  #rm -rf ~/.local/share/apptree/
  if [ \$(command -v xz) ]; then
    #URL=https://raw.githubusercontent.com/alphahlee/workflow/master/cmd/cli/apptree-1.0.tar.gz
    #URL=https://raw.githubusercontent.com/alphahlee/workflow/master/cmd/cli/apptree
    URL="\$URL"
    #TAR_ARGS="xJ"
    TAR_ARGS="-C /usr/local/bin xz"
  else
    #URL=https://raw.githubusercontent.com/alphahlee/workflow/master/cmd/cli/apptree-1.0.tar.gz
    #URL=https://raw.githubusercontent.com/alphahlee/workflow/master/cmd/cli/apptree
    URL="\$URL"
    TAR_ARGS="-C /usr/local/bin xz"
  fi
  
  echo "Installing CLI from \$URL"
  if [ \$(command -v curl) ]; then
    #curl "\$URL" | tar "\$TAR_ARGS"
    curl -o apptree "\$URL"
  else
    #wget -O- "\$URL" | tar "\$TAR_ARGS"
    #wget -O "\$URL"
    wget -O apptree "\$URL"
  fi
  
  # change permissions on apptree executable
  chmod 755 /usr/local/bin/apptree

  # download default cert and key
  CERTURL=https://raw.githubusercontent.com/alphahlee/workflow/master/cmd/cli/server.crt
  KEYURL=https://raw.githubusercontent.com/alphahlee/workflow/master/cmd/cli/server.key

  echo "Installing cert from \$CERTURL"
  cd $HOME/.apptree
  if [ \$(command -v curl) ]; then
    curl -o server.crt "\$CERTURL"
  else
    wget -O "\$CERTURL"
  fi

  echo "Installing key from \$KEYURL"
  cd $HOME/.apptree
  if [ \$(command -v curl) ]; then
    curl -o server.key "\$KEYURL"
  else
    wget -O "\$KEYURL"
  fi

  # delete old apptree bin if exists
  #rm -f \$(command -v apptree) || true
  #rm -f /usr/local/bin/apptree

  # change permissions on $HOME/.apptree directory
  chmod -R 766 $HOME/.apptree/

SCRIPT
  # test the CLI
  LOCATION=$(command -v apptree)
  echo "apptree installed to $LOCATION"
  apptree --version
}
