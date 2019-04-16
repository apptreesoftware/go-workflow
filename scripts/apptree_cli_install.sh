#!/bin/bash

  #unset INSTALL_TYPE
  #unset INSTALL_DIR
  #unset ENGINE_PORT
  #unset STEP_PORT
  #unset CURRENT_DIR
  #unset CERTURL
  #unset KEYURL
    
  # Set Installation Variables
  CURRENT_DIR=$(pwd)
  DARWINURL="https://storage.googleapis.com/apptreeworkflow/binaries/apptree_darwin_amd64"
  LINUXURL="https://storage.googleapis.com/apptreeworkflow/binaries/apptree_linux_amd64"
  CERTURL="https://s3.amazonaws.com/apptree-binaries/server.crt"
  KEYURL="https://s3.amazonaws.com/apptree-binaries/server.key"
  LOCATION="/usr/local/bin/apptree"

  echo '1) Enter the location for the .apptree folder: '
  read INSTALL_DIR
  echo "You entered $INSTALL_DIR for the .apptree folder."
  echo "......"
  APPTREE_PARAM_FILE=$CURRENT_DIR/apptree_install_env_$ENGINE_PORT-$STEP_PORT.sh
  LOGFILE=$CURRENT_DIR/apptree_engine_$ENGINE_PORT-$STEP_PORT-install.log
  APPTREE_DIR=$INSTALL_DIR/.apptree
  INSTALL_USER=`id -un`
  INSTALL_USER_GROUP=`id -gn`
  echo INSTALL_USER $INSTALL_USER
  echo INSTALL_USER_GROUP $INSTALL_USER_GROUP
  echo INSTALL_DIR $INSTALL_DIR
  #echo ENGINE_PORT $ENGINE_PORT
  #echo STEP_PORT $STEP_PORT
  echo LOCATION $LOCATION
  #ENGINE_INSTALL_CMD="apptree install engine --host https://io.apptreesoftware.com --home \$INSTALL_DIR --port \$ENGINE_PORT --step_port \$STEP_PORT -f \$LOCATION"
  #echo ENGINE_INSTALL_CMD $ENGINE_INSTALL_CMD
  CHOWN_CMD="chown -R \$INSTALL_USER:\$INSTALL_USER_GROUP \$INSTALL_DIR"
  echo CHOWN_CMD $CHOWN_CMD
  echo "APPTREE_PARAM_FILE: $APPTREE_PARAM_FILE"
  echo "......"
  echo "Generating Installation parameter file $APPTREE_PARAM_FILE."
  #echo "#!/bin/bash" > $APPTREE_PARAM_FILE
  echo "export INSTALL_TYPE=$INSTALL_TYPE" >> $APPTREE_PARAM_FILE
  echo "export CURRENT_DIR=$CURRENT_DIR" > $APPTREE_PARAM_FILE
  echo "export INSTALL_DIR=$INSTALL_DIR" >> $APPTREE_PARAM_FILE
  echo "export APPTREE_DIR=$INSTALL_DIR/.apptree" >> $APPTREE_PARAM_FILE
  echo "export LOGFILE=$LOGFILE" >> $APPTREE_PARAM_FILE
  echo "export LINUXURL=$LINUXURL" >> $APPTREE_PARAM_FILE
  echo "export DARWINURL=$DARWINURL" >> $APPTREE_PARAM_FILE
  echo "export CERTURL=$CERTURL" >> $APPTREE_PARAM_FILE
  echo "export KEYURL=$KEYURL" >> $APPTREE_PARAM_FILE
  echo "export INSTALL_USER=$INSTALL_USER" >> $APPTREE_PARAM_FILE
  echo "export INSTALL_USER_GROUP=$INSTALL_USER_GROUP" >> $APPTREE_PARAM_FILE
  #echo "export ENGINE_INSTALL_CMD=\"$ENGINE_INSTALL_CMD\"" >> $APPTREE_PARAM_FILE
  echo "export CHOWN_CMD=\"$CHOWN_CMD\"" >> $APPTREE_PARAM_FILE
  chmod 755 $APPTREE_PARAM_FILE
  
  date > $LOGFILE
  
  echo Checking $APPTREE_PARAM_FILE
  source $APPTREE_PARAM_FILE
  
  echo Installation Directory: $INSTALL_DIR >> $LOGFILE
  #echo Engine Port: $ENGINE_PORT >> $LOGFILE
  #echo Step Port: $STEP_PORT >> $LOGFILE
  echo Working Directory: $CURRENT_DIR >> $LOGFILE
  
  if [ -z "$INSTALL_DIR" ]; then
    echo "You did not specify a INSTALL_DIR, please restart the installation and enter a valid directory for installation."
	#mkdir $INSTALL_DIR
    exit 1
  fi

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

  echo Sourcing $APPTREE_PARAM_FILE
  source $APPTREE_PARAM_FILE
  
  #echo PATH=$PATH >> $LOGFILE

  if [ "\$(uname)" == "Darwin" ]; then
    OS=darwin
    #URL="https://storage.googleapis.com/apptreeworkflow/binaries/apptree_darwin_amd64"
        URL=$DARWINURL
    echo "OS is darwin. Download URL is \$URL" >> $LOGFILE
        echo "......"
  elif [ "\$(expr substr \$(uname -s) 1 5)" == "Linux" ]; then
    OS=linux
    #URL="https://storage.googleapis.com/apptreeworkflow/binaries/apptree_linux_amd64"
        URL=$LINUXURL
    echo "OS is linux. Download URL is \$URL" >> $LOGFILE
        echo "......"
    echo Setting PATH
    export PATH=$PATH:/usr/local/bin
  else
    echoerr "This installer is only supported on Linux and MacOS"
    exit 1
  fi

  #echo Setting PATH
  #export PATH=$PATH:/usr/local/bin
  
  if [ ! -d $APPTREE_DIR ]; then
    mkdir -p $APPTREE_DIR
    chmod -R 777 $APPTREE_DIR
  fi

  if [[ ! ":\$PATH:" == *":/usr/local/bin:"* ]]; then
    echo ":\$PATH:"
  #if [[ ! ":$PATH:" == *":/usr/local/bin:"* ]]; then
    echo ":$PATH:"
    echoerr "Your path is missing /usr/local/bin, adding this to your PATH for the installer."
    export PATH=$PATH:/usr/local/bin
    #exit 1
  fi
  
  if [ -z "$INSTALL_DIR" ]; then
    echoerr "You did not specify a INSTALL_DIR, please Control+C and restart the installation."
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
  
  cd /usr/local/bin
  rm -rf apptree
  #rm -rf ~/.local/share/apptree/
  if [ \$(command -v xz) ]; then
    URL="\$URL"
    #TAR_ARGS="xJ"
    TAR_ARGS="-C /usr/local/bin xz"
  else
    URL="\$URL"
    TAR_ARGS="-C /usr/local/bin xz"
  fi
  
  echo "Installing CLI from \$URL"
  if [ \$(command -v curl) ]; then
    curl -o apptree "\$URL"
	echo "Completed CLI Download"
	echo "......"
  else
    wget -O apptree "\$URL"
	echo "Completed CLI Download"
	echo "......"
  fi

  chmod 755 /usr/local/bin/apptree

  echo "......"
  echo AppTree cache, cert, key files will be installed under $INSTALL_DIR/.apptree
  echo "......"
  #echo Creating "$INSTALL_DIR/.apptree" directory
  #mkdir -p $INSTALL_DIR/.apptree
  
  echo "Installing cert from \$CERTURL"
  cd $INSTALL_DIR/.apptree
  if [ \$(command -v curl) ]; then
    curl -o server.crt "\$CERTURL"
	echo "Completed cert Download"
	echo "......"
  else
    wget -O "\$CERTURL"
	echo "Completed cert Download"
	echo "......"
  fi
  echo cert installed under $INSTALL_DIR/.apptree
  echo Cert installed under $INSTALL_DIR/.apptree >> $LOGFILE
  echo "......"

  echo "Installing key from \$KEYURL"
  cd $INSTALL_DIR/.apptree
  if [ \$(command -v curl) ]; then
    curl -o server.key "\$KEYURL"
	echo "Completed key Download"
	echo "......"
  else
    wget -O "\$KEYURL"
	echo "Completed key Download"
	echo "......"
  fi
  echo key installed under $INSTALL_DIR/.apptree
  echo Key installed under $INSTALL_DIR/.apptree >> $LOGFILE
  echo "......"

  echo Changing ownership on $INSTALL_DIR
  echo Changing ownership on $INSTALL_DIR >> $LOGFILE
  #echo $CHOWN_CMD
  echo chown -R $INSTALL_USER:$INSTALL_USER_GROUP $INSTALL_DIR
  #$CHOWN_CMD >> $LOGFILE
  chown -R $INSTALL_USER:$INSTALL_USER_GROUP $INSTALL_DIR >> $LOGFILE
  echo "......"

  echo Changing permissions on $INSTALL_DIR/.apptree
  echo Changing permissions on $INSTALL_DIR/.apptree >> $LOGFILE
  chmod -R 766 $INSTALL_DIR/.apptree/
  echo "......"

SCRIPT
  # test the CLI
  LOCATION=$(command -v apptree)
  echo "apptree installed to $LOCATION"
  apptree --version
  apptree --version >> $LOGFILE
  echo "apptree installation complete!"
  echo "apptree installation complete!" >> $LOGFILE
  date >> $LOGFILE
}
