#!/bin/bash

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

PORT=3010

echo -e "${YELLOW}🔧 Starting Fiber Boilerplate Setup...${NC}\n"

# Step 1: Check and kill existing process on port 3010
echo -e "${BLUE}🔍 Checking if port ${PORT} is already in use...${NC}"
PID=$(lsof -t -i:$PORT)

if [ ! -z "$PID" ]; then
    echo -e "${YELLOW}⚠️  Process found on port ${PORT} (PID: $PID)${NC}"
    echo -e "${YELLOW}🔪 Killing existing process...${NC}"
    kill -9 $PID
    sleep 1
    echo -e "${GREEN}✅ Process killed successfully!${NC}\n"
else
    echo -e "${GREEN}✅ Port ${PORT} is free!${NC}\n"
fi

# Step 2: Generate Swagger Documentation
echo -e "${YELLOW}📝 Generating Swagger documentation...${NC}"
$(go env GOPATH)/bin/swag init -g cmd/main.go -o docs --parseDependency --parseInternal 2>&1 | grep -v "failed to evaluate const"

# Check if swagger generation was successful
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Swagger docs generated successfully!${NC}\n"
    
    # Step 3: Start the server
    echo -e "${YELLOW}🚀 Starting Fiber server...${NC}"
    echo -e "${GREEN}📍 Swagger UI: http://localhost:3010/swagger/index.html${NC}\n"
    
    go run cmd/main.go
else
    echo -e "${RED}❌ Failed to generate Swagger documentation!${NC}"
    echo -e "${RED}Please check if swag is installed: go install github.com/swaggo/swag/cmd/swag@latest${NC}"
    exit 1
fi
