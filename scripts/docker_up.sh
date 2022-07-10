#! /bin/bash

set -e

export APP_MODE=$1

reset=`tput sgr0`
red=`tput setaf 1`
green=`tput setaf 2`
yellow=`tput setaf 3`
blue=`tput setaf 4`

# Finding configuration and docker-compose files for chosen environment.
if [[ -f "$(pwd)/config/${APP_MODE}/config.yaml" && -f "$(pwd)/config/${APP_MODE}/docker-compose.yaml" ]]; then
    composeFile="$(pwd)/config/${APP_MODE}/docker-compose.yaml"

    echo "${green}Application is running on ${blue}${APP_MODE}${green} environment.${reset}"
else
    echo "${red}Failed to find configuration files for ${yellow}${APP_MODE}${red} environment!${reset}"
    echo "${red}You're missing ${yellow}config.yaml${reset}${red} or/and ${yellow}docker-compose.yaml${reset}${red} under path ${yellow}config/${APP_MODE}/${reset}"
    exit;
fi

{
    DOCKER_BUILDKIT=1 docker-compose \
        --file $composeFile \
        --project-name restaurant_gateway \
        up -d --build &&
        echo "${green}Your containers are up and running.${reset}"
} || {
    echo "${red}Something went wrong!${reset}"
}