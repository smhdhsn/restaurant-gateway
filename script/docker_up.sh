#! /bin/bash

reset=`tput sgr0`
red=`tput setaf 1`
green=`tput setaf 2`
yellow=`tput setaf 3`
blue=`tput setaf 4`

# Finding configuration and docker-compose files for chosen environment.
if [ $APP_MODE ]; then
    if [[ -f "$(pwd)/config/${APP_MODE}/config.yaml" && -f "$(pwd)/config/${APP_MODE}/docker-compose.yaml" ]]; then
        composeFile="$(pwd)/config/${APP_MODE}/docker-compose.yaml"

        echo "${green}Application is running on ${blue}${APP_MODE}${green} environment.${reset}"
    else
        echo "${red}Failed to find configuration files for ${yellow}${APP_MODE}${red} environment!${reset}"
        echo "${red}You're missing ${yellow}config.yaml${reset}${red} or/and ${yellow}docker-compose.yaml${reset}${red} under path ${yellow}config/${APP_MODE}/${reset}"
        exit;
    fi
else
    if [[ -f "$(pwd)/config/local/config.yaml" && -f "$(pwd)/config/local/docker-compose.yaml" ]]; then
        export APP_MODE="local"
        composeFile="$(pwd)/config/${APP_MODE}/docker-compose.yaml"

        echo "${green}Application is running on ${blue}${APP_MODE}${green} environment.${reset}"
    else
        echo "${red}Failed to find configuration files for any environment!${reset}"
        echo "${red}You're missing ${yellow}config.yaml${reset}${red} or/and ${yellow}docker-compose.yaml${reset}${red} under path ${yellow}$(pwd)/config/APP_MODE/${reset}"
        exit;
    fi
fi

DOCKER_BUILDKIT=1 docker-compose \
    --file $composeFile \
    --project-name restaurant_gateway \
    up -d --build
