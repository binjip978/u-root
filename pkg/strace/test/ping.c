#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <sys/socket.h>
#include <netinet/ip_icmp.h>
#include <netinet/in.h>
#include <sys/time.h>

// ICMP header structure
struct icmpheader {
    unsigned char type;
    unsigned char code;
    unsigned short checksum;
    unsigned short id;
    unsigned short sequence;
};

// IP header structure
struct ipheader {
    unsigned char iph_ihl:4, iph_ver:4;
    unsigned char iph_tos;
    unsigned short iph_len;
    unsigned short iph_ident;
    unsigned short iph_flag:3, iph_offset:13;
    unsigned char iph_ttl;
    unsigned char iph_protocol;
    unsigned short iph_chksum;
    struct in_addr iph_sourceip;
    struct in_addr iph_destip;
};

// Calculate checksum
unsigned short calculate_checksum(void *b, int len) {
    unsigned short *buf = b;
    unsigned int sum = 0;
    unsigned short result;

    for (sum = 0; len > 1; len -= 2) {
        sum += *buf++;
    }
    if (len == 1) {
        sum += *(unsigned char *)buf;
    }
    sum = (sum >> 16) + (sum & 0xFFFF);
    sum += (sum >> 16);
    result = ~sum;
    return result;
}

int main() {
    struct sockaddr_in addr;
    int sockfd;
    char packet[64];
    struct ipheader *ip = (struct ipheader *)packet;
    struct icmpheader *icmp = (struct icmpheader *)(packet + sizeof(struct ipheader));
    struct timeval start, end;
    double elapsed;

    // Create raw socket
    sockfd = socket(AF_INET, SOCK_RAW, IPPROTO_ICMP);
    if (sockfd < 0) {
        exit(EXIT_FAILURE);
    }

    // Fill in the sockaddr_in structure
    addr.sin_family = AF_INET;
    addr.sin_addr.s_addr = inet_addr("127.0.0.1");

    // Fill in the ICMP header
    icmp->type = ICMP_ECHO;
    icmp->code = 0;
    icmp->checksum = 0;
    icmp->id = getpid();
    icmp->sequence = 0;

    // Calculate ICMP checksum
    icmp->checksum = calculate_checksum((unsigned short *)icmp, sizeof(struct icmpheader));

    // Send the packet
    gettimeofday(&start, NULL);
    if (sendto(sockfd, packet, sizeof(struct icmpheader), 0, (struct sockaddr *)&addr, sizeof(addr)) < 0) {
        exit(EXIT_FAILURE);
    }

    // Wait for the response
    if (recv(sockfd, packet, sizeof(packet), 0) < 0) {
        exit(EXIT_FAILURE);
    }
    gettimeofday(&end, NULL);

    // Calculate elapsed time
    elapsed = (end.tv_sec - start.tv_sec) * 1000.0;
    elapsed += (end.tv_usec - start.tv_usec) / 1000.0;

    printf("Ping received in %.2f ms\n", elapsed);

    close(sockfd);
}

