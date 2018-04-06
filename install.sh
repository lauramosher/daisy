#!/bin/sh

printf "\033[0;36m\xF0\x9F\x94\x91  Installing Daisy\033[m\n"

printf "   Fetching latest release binary..."
cd /usr/local/bin && { curl -OLs https://github.com/lauramosher/daisy/releases/download/v0.2.0/day ; cd $OLDPWD ; }

if [ -e /usr/local/bin/day ]
then
  printf "\033[0;32m\t\t\t\t\tDone!\033[m\n"
else
  printf "\n   Unable to download binary. Please see manual install instructions.\n"
  exit
fi

echo "   Configuring Slack Integation..."

if [[ -z $DAISY_SLACK_USER_TOKEN ]]
then
  printf "\033[1;35m"
  printf "\xE2\x9A\xA0  Daisy needs to permission to post as you in Slack\n"
  printf "\033[1;36m"

  printf "   --------------------------------------------------------------------------\n"
  printf "   Daisy requires a User Token for the Slack integration to function.\n\n"
  printf "   Please follow the instructions to authorize Daisy on your Slack workspace:\n"
  printf "     https://daisy-slack.herokuapp.com/\n"
  printf "   --------------------------------------------------------------------------\n\n"
  printf "\033[m"

  printf "\033[0;35m   Do you have your User Access Token? (Yn) \033[m"
  read hastoken

  if [ -z $hastoken ] || [ $hastoken != "Y" ]; then
    printf "   Opening Slack authorization..."

    open "https://slack.com/oauth/authorize?client_id=2156327032.337868849991&scope=emoji:read,users.profile:write,users:write,chat:write:user"

    printf "\033[0;32m\t\t\t\t\tDone!\033[m\n\n"
    printf "   Authorize slack...\n"
  fi

  printf "\xE2\x9A\xA1  Enter Access Token: "
  read token
fi

if [[ -z $DAISY_SLACK_CHANNEL ]]
then
  printf "\xE2\x9A\xA1  What channel do you want daisy to post to? "
  read channel
fi

printf "   Adding configuration to your bash/zsh profile..."
if [ -e $HOME/.zshrc ] && ! grep -q "Daisy" $HOME/.zshrc
then
  echo "#######" >> $HOME/.zshrc
  echo "# Daisy\n" >> $HOME/.zshrc
  echo "export DAISY_SLACK_USER_TOKEN=$token" >> $HOME/.zshrc
  echo "export DAISY_SLACK_CHANNEL=$channel" >> $HOME/.zshrc
elif [ -e $HOME/.bash_profile ] && ! grep -q "Daisy" $HOME/.bash_profile
then
  echo "#######" >> $HOME/.bash_profile
  echo "# Daisy\n" >> $HOME/.bash_profile
  echo "export DAISY_SLACK_USER_TOKEN=$token" >> $HOME/.bash_profile
  echo "export DAISY_SLACK_CHANNEL=$channel" >> $HOME/.bash_profile
elif [ -e $HOME/.bashrc ] && ! grep -q "Daisy" $HOME/.bashrc
then
  echo "#######" >> $HOME/.bashrc
  echo "# Daisy\n" >> $HOME/.bashrc
  echo "export DAISY_SLACK_USER_TOKEN=$token" >> $HOME/.bashrc
  echo "export DAISY_SLACK_CHANNEL=$channel" >> $HOME/.bashrc
fi

printf "\033[0;32m\t\t\tDone!\033[m\n\n"
printf "\033[1;31m\xE2\x9D\x97  Please restart your terminal for the changes to take effect!\033[m\n\n"

printf "\xF0\x9F\x99\x87 Thank you for choosing Daisy!\n\n"
