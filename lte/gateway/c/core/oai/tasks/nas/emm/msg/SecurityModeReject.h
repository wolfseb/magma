/*
 * Licensed to the OpenAirInterface (OAI) Software Alliance under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The OpenAirInterface Software Alliance licenses this file to You under
 * the terms found in the LICENSE file in the root of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *-------------------------------------------------------------------------------
 * For more information about the OpenAirInterface (OAI) Software Alliance:
 *      contact@openairinterface.org
 */

#ifndef FILE_SECURITY_MODE_REJECT_SEEN
#define FILE_SECURITY_MODE_REJECT_SEEN
#include <stdint.h>

#include "lte/gateway/c/core/oai/tasks/nas/ies/SecurityHeaderType.h"
#include "lte/gateway/c/core/oai/tasks/nas/ies/MessageType.h"
#include "lte/gateway/c/core/oai/tasks/nas/ies/EmmCause.h"
#include "lte/gateway/c/core/oai/lib/3gpp/3gpp_23.003.h"
#include "lte/gateway/c/core/oai/lib/3gpp/3gpp_24.007.h"
#include "lte/gateway/c/core/oai/lib/3gpp/3gpp_24.008.h"

/* Minimum length macro. Formed by minimum length of each mandatory field */
#define SECURITY_MODE_REJECT_MINIMUM_LENGTH (EMM_CAUSE_MINIMUM_LENGTH)

/* Maximum length macro. Formed by maximum length of each field */
#define SECURITY_MODE_REJECT_MAXIMUM_LENGTH (EMM_CAUSE_MAXIMUM_LENGTH)

/*
 * Message name: Security mode reject
 * Description: This message is sent by the UE to the network to indicate that
 * the corresponding security mode command has been rejected. See
 * table 8.2.22.1. Significance: dual Direction: UE to network
 */

typedef struct security_mode_reject_msg_tag {
  /* Mandatory fields */
  eps_protocol_discriminator_t protocoldiscriminator : 4;
  security_header_type_t securityheadertype : 4;
  message_type_t messagetype;
  emm_cause_t emmcause;
} security_mode_reject_msg;

int decode_security_mode_reject(security_mode_reject_msg* securitymodereject,
                                uint8_t* buffer, uint32_t len);

int encode_security_mode_reject(security_mode_reject_msg* securitymodereject,
                                uint8_t* buffer, uint32_t len);

#endif /* ! defined(FILE_SECURITY_MODE_REJECT_SEEN) */
